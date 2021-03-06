package i18n

import "golang.org/x/text/language"

type translationPair struct {
	key string
	msg string
}

var translationsSet = map[language.Tag][]translationPair{
	language.English: {
		{"title", "Private and secure notes - send your secrets safely."},
		{"header", "Private secure notes"},
		{"description", "Highly secure message encryption and open source tool that auto-destruct."},
		{"enterTextMessage", "Enter text message to encrypt"},
		{"secureButton", "Encrypt message"},
		{"copyLink", "Copy link and send it to a friend. Message will be deleted after being read or after 4 weeks when not read."},
		{"copyLinkButton", "Copy link"},
		{"newMessageButton", "New message"},
		{"decodedMessage", "Decrypted message"},
		{"messageRead", "Message was already read, deleted or link is corrupted"},
		{"readMessageButton", "Read message"},
		{"infoHeader", "info about"},
		{"info", "This tool was built with care and respect to your privacy. " +
			"Tool uses various method of encryption to ensure maximum privacy. To increase security feel free to use password. " +
			"Tool is Open Source and code is publicly accessible. You can see how it works on"},
		{"info1", "If you want to contact us please send an"},
		{"info2", "email"},
		{"info3", "All translation initiatives are welcome."},
		{"linkCorrupted", "Link is corrupted"},
		{"generalError", "Something went wrong. Try again later."},
		{"encryptNetworkError", "Something went wrong. Cannot save the message. Please try again."},
		{"decryptNetworkError", "Something went wrong. Cannot load the message. Please try again."},
		{"password", "Password"},
		{"passwordEncryptPlaceholder", "Optional password to increase security"},
		{"passwordDecryptPlaceholder", "enter password"},
		{"linkIsCorrupted", "Link is corrupted"},
		{"ieEncryptWarning", "Internet Explorer detected. Encryption may take a few seconds. Please be patient."},
		{"ieDecryptWarning", "Internet Explorer detected. Decryption may take a few seconds. Please be patient."},
	},
	language.Polish: {
		{"title", "Prywatne bezpieczne wiadomości"},
		{"header", "Prywatne wiadomości"},
		{"description", "Bezpieczne szyfrowane wiadomości"},
		{"enterTextMessage", "Wpisz wiadomość do zaszyfrowania"},
		{"secureButton", "Szyfruj wiadomość"},
		{"copyLink", "Skopiuj link i prześlij do przyjaciela. Wiadomość będzie skasowana natychmiast po odczytaniu lub po 4 tygodniach."},
		{"copyLinkButton", "Skopiuj link"},
		{"newMessageButton", "Nowa wiadomość"},
		{"decodedMessage", "Odszyfrowana wiadomość"},
		{"messageRead", "Wiadomość odczytana, przeterminowana lub link jest błędny"},
		{"readMessageButton", "Odszyfruj wiadomość"},
		{"infoHeader", "opis"},
		{"info", "Narzędzie używa różnych metod szyfrowania, aby zapewnić maksymalne bezpieczeństwo. " +
			"Bez posiadania linku nie ma możliwości odszyfrowania wiadomości. Użyj hasła aby dodatkowo zwiększyć bezpieczeństwo. " +
			"Kod źródłowy narzędzia jest otwarty i możesz go obejrzeć w serwisie"},
		{"info1", "Jeśli chcesz się z nami skontaktować wyślij nam"},
		{"info2", "wiadomość"},
		{"info3", ""},
		{"linkCorrupted", "Link uszkodzony"},
		{"generalError", "Coś poszło nie tak. Spróbuj ponownie za jakiś czas."},
		{"encryptNetworkError", "Coś poszło nie tak. Nie mogę zapisać wiadomości. Spróbuj ponownie."},
		{"decryptNetworkError", "Coś poszło nie tak. Nie mogę odczytać wiadomości. Spróbuj ponownie."},
		{"password", "Hasło"},
		{"passwordEncryptPlaceholder", "Możesz użyć hasła aby zwiększyć bezpieczeństwo"},
		{"passwordDecryptPlaceholder", "wprowadź hasło"},
		{"linkIsCorrupted", "Link jest uszkodzony"},
		{"ieEncryptWarning", "Używasz Internet Explorera. Szyfrowanie może potrwać parę sekund. Proszę o cierpliwość."},
		{"ieDecryptWarning", "Używasz Internet Explorera. Odszyfrowanie może potrwać parę sekund. Proszę o cierpliwość."},
	},
}
