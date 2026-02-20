package i18n_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/OpenSlides/openslides-projector-service/pkg/i18n"
	"golang.org/x/text/language"
)

func initHelloWorldFile(t *testing.T) {
	tmp := t.TempDir()

	poDir := filepath.Join(tmp, "locale", "de")
	if err := os.MkdirAll(poDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	poContent := `msgid ""
msgstr ""
"Content-Type: text/plain; charset=UTF-8\n"
"Language: de\n"

msgid "Hello"
msgstr "Hallo"
`
	if err := os.WriteFile(filepath.Join(poDir, "default.po"), []byte(poContent), 0o644); err != nil {
		t.Fatalf("write po: %v", err)
	}

	// Switch to temp dir so NewLocale("locale", ...) resolves correctly.
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	if err := os.Chdir(tmp); err != nil {
		t.Fatalf("chdir: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Chdir(oldWd)
	})
}

func TestProjectorLocaleGetFallback(t *testing.T) {
	locale := i18n.NewLocale(language.English)
	const input = "Hello world"

	got := locale.Get(input)
	if got != input {
		t.Fatalf("expected fallback to input string, got %q", got)
	}
}

func TestProjectorLocaleGetUsesTranslationFile(t *testing.T) {
	initHelloWorldFile(t)

	locale := i18n.NewLocale(language.German)

	got := locale.Get("Hello")
	if got != "Hallo" {
		t.Errorf("expected translated string, got %q", got)
	}
}

func TestProjectorLocaleCustomTranslation(t *testing.T) {
	initHelloWorldFile(t)

	locale := i18n.NewLocale(language.English)
	locale.SetCustomTranslations(map[string]string{
		"Hello": "nuqneH",
	})

	got := locale.Get("Hello")
	if got != "nuqneH" {
		t.Errorf("expected custom translation, got %q", got)
	}
}

func TestProjectorLocaleCustomTranslationNested(t *testing.T) {
	initHelloWorldFile(t)

	locale := i18n.NewLocale(language.German)
	locale.SetCustomTranslations(map[string]string{
		"Hallo": "nuqneH",
	})

	got := locale.Get("Hello")
	if got != "nuqneH" {
		t.Errorf("expected custom translation, got %q", got)
	}
}
