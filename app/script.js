// scripts.js
document.addEventListener("DOMContentLoaded", function () {
    // Function to save form data to localStorage
    function saveFormData() {
        const formElements = document.querySelectorAll("#twamp-form input");
        formElements.forEach(element => {
            localStorage.setItem(element.id, element.value);
        });
    }

    // Function to load form data from localStorage
    function loadFormData() {
        const formElements = document.querySelectorAll("#twamp-form input");
        formElements.forEach(element => {
            if (localStorage.getItem(element.id)) {
                element.value = localStorage.getItem(element.id);
            }
        });
    }

    // Load form data on page load
    loadFormData();

    // Save form data on input change
    const formElements = document.querySelectorAll("#twamp-form input");
    formElements.forEach(element => {
        element.addEventListener("input", saveFormData);
    });

    // Clear form data from localStorage on form submission
    document.getElementById('twamp-form').addEventListener('submit', function() {
        const formElements = document.querySelectorAll("#twamp-form input");
        formElements.forEach(element => {
            localStorage.removeItem(element.id);
        });
    });
});
