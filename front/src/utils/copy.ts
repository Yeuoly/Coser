export const copyToClipboard = (text: string) => {
    if (navigator.clipboard) {
        return navigator.clipboard.writeText(text);
    } else {
        var textArea = document.createElement("textarea");
        textArea.value = text;
        textArea.style.position = "fixed";
        document.body.appendChild(textArea);
        textArea.focus();
        textArea.select();

        try {
            document.execCommand('copy');
        } catch (err) { }

        document.body.removeChild(textArea);
    }
}