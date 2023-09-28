// On page load, check user's preference
window.onload = function() {
    const theme = localStorage.getItem('theme');
    const modeToggleElement = document.getElementById('modeToggle');

    if (theme === 'dark') {
        document.body.classList.add('dark-mode');
        modeToggleElement.textContent = '☀'; // Sun symbol for dark mode
    } else {
        document.body.classList.remove('dark-mode');
        modeToggleElement.textContent = '🌙'; // Moon symbol for light mode
    }
};

document.getElementById('modeToggle').addEventListener('click', function() {
    const modeToggleElement = document.getElementById('modeToggle');
    document.body.classList.toggle('dark-mode');
    
    if (document.body.classList.contains('dark-mode')) {
        localStorage.setItem('theme', 'dark');
        modeToggleElement.textContent = '☀'; // Sun symbol for dark mode
    } else {
        localStorage.setItem('theme', 'light');
        modeToggleElement.textContent = '🌙'; // Moon symbol for light mode
    }
});

