function colorMode() {
    const storedTheme = localStorage.getItem("theme");
    
    const systemPrefersDark = window.matchMedia("(prefers-color-scheme: dark)").matches;
    const useDark = storedTheme === "dark" || (!storedTheme && systemPrefersDark);

    return {
        darkMode: useDark,
        toggleDarkMode() {
            document.documentElement.classList.toggle('dark');
            this.darkMode = !this.darkMode;

            localStorage.theme = this.darkMode ? 'dark' : 'light';
        }
    }
}