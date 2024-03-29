// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/dashotv/tmdb/openapi/utils"
	"net/http"
)

type WatchProviderTvListRequest struct {
	Language    *string `default:"en-US" queryParam:"style=form,explode=true,name=language"`
	WatchRegion *string `queryParam:"style=form,explode=true,name=watch_region"`
}

func (w WatchProviderTvListRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WatchProviderTvListRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WatchProviderTvListRequest) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

func (o *WatchProviderTvListRequest) GetWatchRegion() *string {
	if o == nil {
		return nil
	}
	return o.WatchRegion
}

type WatchProviderTvList200ApplicationJSONResultsDisplayPriorities struct {
	Ae *int64 `default:"0" json:"AE"`
	Ar *int64 `default:"0" json:"AR"`
	At *int64 `default:"0" json:"AT"`
	Au *int64 `default:"0" json:"AU"`
	Be *int64 `default:"0" json:"BE"`
	Bg *int64 `default:"0" json:"BG"`
	Bo *int64 `default:"0" json:"BO"`
	Br *int64 `default:"0" json:"BR"`
	Ca *int64 `default:"0" json:"CA"`
	Ch *int64 `default:"0" json:"CH"`
	Cl *int64 `default:"0" json:"CL"`
	Co *int64 `default:"0" json:"CO"`
	Cr *int64 `default:"0" json:"CR"`
	Cv *int64 `default:"0" json:"CV"`
	Cz *int64 `default:"0" json:"CZ"`
	De *int64 `default:"0" json:"DE"`
	Dk *int64 `default:"0" json:"DK"`
	Ec *int64 `default:"0" json:"EC"`
	Ee *int64 `default:"0" json:"EE"`
	Eg *int64 `default:"0" json:"EG"`
	Es *int64 `default:"0" json:"ES"`
	Fi *int64 `default:"0" json:"FI"`
	Fr *int64 `default:"0" json:"FR"`
	Gb *int64 `default:"0" json:"GB"`
	Gh *int64 `default:"0" json:"GH"`
	Gr *int64 `default:"0" json:"GR"`
	Gt *int64 `default:"0" json:"GT"`
	Hk *int64 `default:"0" json:"HK"`
	Hn *int64 `default:"0" json:"HN"`
	Hu *int64 `default:"0" json:"HU"`
	ID *int64 `default:"0" json:"ID"`
	Ie *int64 `default:"0" json:"IE"`
	Il *int64 `default:"0" json:"IL"`
	In *int64 `default:"0" json:"IN"`
	It *int64 `default:"0" json:"IT"`
	Jp *int64 `default:"0" json:"JP"`
	Lt *int64 `default:"0" json:"LT"`
	Lv *int64 `default:"0" json:"LV"`
	Mu *int64 `default:"0" json:"MU"`
	Mx *int64 `default:"0" json:"MX"`
	My *int64 `default:"0" json:"MY"`
	Mz *int64 `default:"0" json:"MZ"`
	Nl *int64 `default:"0" json:"NL"`
	No *int64 `default:"0" json:"NO"`
	Nz *int64 `default:"0" json:"NZ"`
	Pe *int64 `default:"0" json:"PE"`
	Ph *int64 `default:"0" json:"PH"`
	Pl *int64 `default:"0" json:"PL"`
	Pt *int64 `default:"0" json:"PT"`
	Py *int64 `default:"0" json:"PY"`
	Ru *int64 `default:"0" json:"RU"`
	Sa *int64 `default:"0" json:"SA"`
	Se *int64 `default:"0" json:"SE"`
	Sg *int64 `default:"0" json:"SG"`
	Si *int64 `default:"0" json:"SI"`
	Sk *int64 `default:"0" json:"SK"`
	Th *int64 `default:"0" json:"TH"`
	Tr *int64 `default:"0" json:"TR"`
	Tw *int64 `default:"0" json:"TW"`
	Ug *int64 `default:"0" json:"UG"`
	Us *int64 `default:"0" json:"US"`
	Ve *int64 `default:"0" json:"VE"`
	Za *int64 `default:"0" json:"ZA"`
}

func (w WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetAe() *int64 {
	if o == nil {
		return nil
	}
	return o.Ae
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetAr() *int64 {
	if o == nil {
		return nil
	}
	return o.Ar
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetAt() *int64 {
	if o == nil {
		return nil
	}
	return o.At
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetAu() *int64 {
	if o == nil {
		return nil
	}
	return o.Au
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetBe() *int64 {
	if o == nil {
		return nil
	}
	return o.Be
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetBg() *int64 {
	if o == nil {
		return nil
	}
	return o.Bg
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetBo() *int64 {
	if o == nil {
		return nil
	}
	return o.Bo
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetBr() *int64 {
	if o == nil {
		return nil
	}
	return o.Br
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetCa() *int64 {
	if o == nil {
		return nil
	}
	return o.Ca
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetCh() *int64 {
	if o == nil {
		return nil
	}
	return o.Ch
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetCl() *int64 {
	if o == nil {
		return nil
	}
	return o.Cl
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetCo() *int64 {
	if o == nil {
		return nil
	}
	return o.Co
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetCr() *int64 {
	if o == nil {
		return nil
	}
	return o.Cr
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetCv() *int64 {
	if o == nil {
		return nil
	}
	return o.Cv
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetCz() *int64 {
	if o == nil {
		return nil
	}
	return o.Cz
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetDe() *int64 {
	if o == nil {
		return nil
	}
	return o.De
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetDk() *int64 {
	if o == nil {
		return nil
	}
	return o.Dk
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetEc() *int64 {
	if o == nil {
		return nil
	}
	return o.Ec
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetEe() *int64 {
	if o == nil {
		return nil
	}
	return o.Ee
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetEg() *int64 {
	if o == nil {
		return nil
	}
	return o.Eg
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetEs() *int64 {
	if o == nil {
		return nil
	}
	return o.Es
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetFi() *int64 {
	if o == nil {
		return nil
	}
	return o.Fi
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetFr() *int64 {
	if o == nil {
		return nil
	}
	return o.Fr
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetGb() *int64 {
	if o == nil {
		return nil
	}
	return o.Gb
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetGh() *int64 {
	if o == nil {
		return nil
	}
	return o.Gh
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetGr() *int64 {
	if o == nil {
		return nil
	}
	return o.Gr
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetGt() *int64 {
	if o == nil {
		return nil
	}
	return o.Gt
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetHk() *int64 {
	if o == nil {
		return nil
	}
	return o.Hk
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetHn() *int64 {
	if o == nil {
		return nil
	}
	return o.Hn
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetHu() *int64 {
	if o == nil {
		return nil
	}
	return o.Hu
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetIe() *int64 {
	if o == nil {
		return nil
	}
	return o.Ie
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetIl() *int64 {
	if o == nil {
		return nil
	}
	return o.Il
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetIn() *int64 {
	if o == nil {
		return nil
	}
	return o.In
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetIt() *int64 {
	if o == nil {
		return nil
	}
	return o.It
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetJp() *int64 {
	if o == nil {
		return nil
	}
	return o.Jp
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetLt() *int64 {
	if o == nil {
		return nil
	}
	return o.Lt
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetLv() *int64 {
	if o == nil {
		return nil
	}
	return o.Lv
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetMu() *int64 {
	if o == nil {
		return nil
	}
	return o.Mu
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetMx() *int64 {
	if o == nil {
		return nil
	}
	return o.Mx
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetMy() *int64 {
	if o == nil {
		return nil
	}
	return o.My
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetMz() *int64 {
	if o == nil {
		return nil
	}
	return o.Mz
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetNl() *int64 {
	if o == nil {
		return nil
	}
	return o.Nl
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetNo() *int64 {
	if o == nil {
		return nil
	}
	return o.No
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetNz() *int64 {
	if o == nil {
		return nil
	}
	return o.Nz
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetPe() *int64 {
	if o == nil {
		return nil
	}
	return o.Pe
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetPh() *int64 {
	if o == nil {
		return nil
	}
	return o.Ph
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetPl() *int64 {
	if o == nil {
		return nil
	}
	return o.Pl
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetPt() *int64 {
	if o == nil {
		return nil
	}
	return o.Pt
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetPy() *int64 {
	if o == nil {
		return nil
	}
	return o.Py
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetRu() *int64 {
	if o == nil {
		return nil
	}
	return o.Ru
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetSa() *int64 {
	if o == nil {
		return nil
	}
	return o.Sa
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetSe() *int64 {
	if o == nil {
		return nil
	}
	return o.Se
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetSg() *int64 {
	if o == nil {
		return nil
	}
	return o.Sg
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetSi() *int64 {
	if o == nil {
		return nil
	}
	return o.Si
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetSk() *int64 {
	if o == nil {
		return nil
	}
	return o.Sk
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetTh() *int64 {
	if o == nil {
		return nil
	}
	return o.Th
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetTr() *int64 {
	if o == nil {
		return nil
	}
	return o.Tr
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetTw() *int64 {
	if o == nil {
		return nil
	}
	return o.Tw
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetUg() *int64 {
	if o == nil {
		return nil
	}
	return o.Ug
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetUs() *int64 {
	if o == nil {
		return nil
	}
	return o.Us
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetVe() *int64 {
	if o == nil {
		return nil
	}
	return o.Ve
}

func (o *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities) GetZa() *int64 {
	if o == nil {
		return nil
	}
	return o.Za
}

type WatchProviderTvList200ApplicationJSONResults struct {
	DisplayPriorities *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities `json:"display_priorities,omitempty"`
	DisplayPriority   *int64                                                         `default:"0" json:"display_priority"`
	LogoPath          *string                                                        `json:"logo_path,omitempty"`
	ProviderID        *int64                                                         `default:"0" json:"provider_id"`
	ProviderName      *string                                                        `json:"provider_name,omitempty"`
}

func (w WatchProviderTvList200ApplicationJSONResults) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WatchProviderTvList200ApplicationJSONResults) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WatchProviderTvList200ApplicationJSONResults) GetDisplayPriorities() *WatchProviderTvList200ApplicationJSONResultsDisplayPriorities {
	if o == nil {
		return nil
	}
	return o.DisplayPriorities
}

func (o *WatchProviderTvList200ApplicationJSONResults) GetDisplayPriority() *int64 {
	if o == nil {
		return nil
	}
	return o.DisplayPriority
}

func (o *WatchProviderTvList200ApplicationJSONResults) GetLogoPath() *string {
	if o == nil {
		return nil
	}
	return o.LogoPath
}

func (o *WatchProviderTvList200ApplicationJSONResults) GetProviderID() *int64 {
	if o == nil {
		return nil
	}
	return o.ProviderID
}

func (o *WatchProviderTvList200ApplicationJSONResults) GetProviderName() *string {
	if o == nil {
		return nil
	}
	return o.ProviderName
}

// WatchProviderTvList200ApplicationJSON - 200
type WatchProviderTvList200ApplicationJSON struct {
	Results []WatchProviderTvList200ApplicationJSONResults `json:"results,omitempty"`
}

func (o *WatchProviderTvList200ApplicationJSON) GetResults() []WatchProviderTvList200ApplicationJSONResults {
	if o == nil {
		return nil
	}
	return o.Results
}

type WatchProviderTvListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// 200
	WatchProviderTvList200ApplicationJSONObject *WatchProviderTvList200ApplicationJSON
}

func (o *WatchProviderTvListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *WatchProviderTvListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *WatchProviderTvListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *WatchProviderTvListResponse) GetWatchProviderTvList200ApplicationJSONObject() *WatchProviderTvList200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.WatchProviderTvList200ApplicationJSONObject
}
