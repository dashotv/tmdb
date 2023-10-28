#!/usr/bin/env ruby

skipped = [
  # these weren't working
  "ConfigurationCountries",
  "ConfigurationJobs",
  "ConfigurationLanguages",
  "ConfigurationPrimaryTranslations",
  "ConfigurationTimezones",
  # not sure why on this one
  "MovieChanges",
  # these are privileged and need user auth, not supported for now
  "AccountAddFavorite",
  "AccountAddToWatchlist",
  "AccountDetails",
  "AccountFavoriteTv",
  "AccountGetFavorites",
  "AccountLists",
  "AccountRatedMovies",
  "AccountRatedTv",
  "AccountRatedTvEpisodes",
  "AccountWatchlistMovies",
  "AccountWatchlistTv",
  "MovieAccountStates",
  "MovieAddRating",
  "MovieDeleteRating",
  "AlternativeNamesCopy",
  "DetailsCopy",
  "MovieLists",
  "ListAddMovie",
  "ListCheckItemStatus",
  "ListClear",
  "ListCreate",
  "ListDelete",
  "ListDetails",
  "ListRemoveMovie",
  "TvEpisodeAccountStates",
  "TvEpisodeAddRating",
  "TvEpisodeDeleteRating",
  "TvSeasonAccountStates",
  "TvSeriesAccountStates",
  "TvSeriesAddRating",
  "TvSeriesDeleteRating",
  # these are auth related for above, not supported for now
  "AuthenticationCreateGuestSession",
  "AuthenticationCreateRequestToken",
  "AuthenticationCreateSession",
  "AuthenticationCreateSessionFromLogin",
  "AuthenticationCreateSessionFromV4Token",
  "AuthenticationDeleteSession",
  "AuthenticationValidateKey",
  "GuestSessionRatedMovies",
  "GuestSessionRatedTv",
  "GuestSessionRatedTvEpisodes",
  # Buggy
  "TvSeriesChanges",
]

FUNC = open('functions.go', 'a+')
TEST = open('functions_test.go', 'a+')

def testVar(n, t)
  "var #{n} #{t} = #{testType(n, t)}"
end

def testType(n, t)
  return 'nil' if t[0] == '*'
  return t.gsub(/^operations\./, 'operations_') if t =~ /^operations\./
  return "#{n}_#{t}"
end

STDIN.each do |line|
  m = line.match(/func \(\w\s\*(\w+)\) (\w+)\(ctx context.Context(, )*([^\)]+)*\) \(\*operations\.(\w+), error\) \{/)
  next unless m
  next if skipped.include?(m[2])
  serv=m[1]
  serv[0]=serv[0].upcase # capitalize will lowercase the rest of the string
  params=""
  testparams=""
  if !m[4].nil?
    params=", " + m[4].split(',').map{|p| p.split(' ')}.map{|p| p[0]}.join(', ')
    testparams=m[4].split(',').map{|p| p.split(' ')}.map{|p| testVar(p[0], p[1])}.join("\t\n")
  end
  FUNC.puts <<-HERE
// #{m[2]} wraps the generated openapi.SDK.#{m[2]} call
func (c *Client) #{m[2]}(#{m[4]}) (*#{m[5]}, error) {
	r, err := c.sdk.#{m[2]}(c.ctx#{params})
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.#{m[2]}200ApplicationJSONObject, nil
}

HERE
  TEST.puts <<-HERE
func TestClient_#{m[2]}(t *testing.T) {
	c := testClient(t)
	#{testparams}
	r, err := c.#{m[2]}(#{params.split(',').drop(1).join(', ')})
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

HERE
end
