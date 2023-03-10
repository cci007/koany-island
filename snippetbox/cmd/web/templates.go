package main

import (
    "html/template"
    "path/filepath"
    "time"
    "cci007/snippetbox/pkg/forms"

    "cci007/snippetbox/pkg/models"


)

// Define a templateData type to act as the holding structure for
// any dynamic data that we want to pass to our HTML templates.
// At the moment it only contains one field, but we'll add more
// to it as the build progresses.
type templateData struct {
	  CSRFToken   string
    CurrentYear int
    Flash       string
    Form        *forms.Form
    IsAuthenticated bool
    Snippet     *models.Snippet
    Snippets    []*models.Snippet
}



func humanDate(t time.Time) string {
    // Return the empty string if time has the zero value.
    if t.IsZero() {
        return ""
    }

    // Convert the time to UTC before formatting it.
    return t.UTC().Format("02 Jan 2006 at 15:04")

}

// Initialize a template.FuncMap object and store it in a global variable. This is
// essentially a string-keyed map which acts as a lookup between the names of our
// custom template functions and the functions themselves.
var functions = template.FuncMap{
    "humanDate": humanDate,
}



func newTemplateCache(dir string) (map[string]*template.Template, error) {
    // Initialize a new map to act as the cache.
    cache := map[string]*template.Template{}

    // Use the filepath.Glob function to get a slice of all filepaths with
    // the extension '.page.tmpl'. This essentially gives us a slice of all the
    // 'page' templates for the application.
    pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
    if err != nil {
        return nil, err
    }

    // Loop through the pages one-by-one.
    for _, page := range pages {
        // Extract the file name (like 'home.page.tmpl') from the full file path
        // and assign it to the name variable.
        name := filepath.Base(page)
        ts, err := template.New(name).Funcs(functions).ParseFiles(page)
        if err != nil {
            return nil, err
        }

        ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
        if err != nil {
            return nil, err
        }

        ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
        if err != nil {
            return nil, err
        }

        cache[name] = ts

    }

    // Return the map.
    return cache, nil
}

