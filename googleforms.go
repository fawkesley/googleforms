package googleforms

import (
	"fmt"
	"net/http"
	"net/url"
)

// Form represents a Google Form that you can post data to.
type Form struct {
	// FormID can be extracted from the "preview form" page URL e.g.
	// https://docs.google.com/forms/d/e/EXTRACT_THIS_FORM_ID/viewform
	FormID string

	// Questions maps a string slug to a Google Forms <input name="?">
	// For example {"email": "entry.123456"}
	Questions map[string]string
}

// Post submits the given answers to the Google Form.
func (f Form) Post(answers map[string]string) error {

	formValues := url.Values{}

	for slug, answer := range answers {
		entryID, ok := f.Questions[slug]
		if !ok {
			return fmt.Errorf("question slug `%s` was not configured at setup", slug)
		}

		formValues[entryID] = []string{answer}
	}

	response, err := http.PostForm(f.postURL(), formValues)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		// bodyBytes, err := io.ReadAll(response.Body)
		// if err == nil {
		// 	fmt.Println(string(bodyBytes))
		// }
		return fmt.Errorf("google forms URL returned HTTP %d", response.StatusCode)
	}
	return nil
}

func (f Form) postURL() string {
	return fmt.Sprintf("https://docs.google.com/forms/u/0/d/e/%s/formResponse", f.FormID)

}
