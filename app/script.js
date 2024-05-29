// scripts.js
document.addEventListener("DOMContentLoaded", function () 
{
    // Function to save form data to localStorage
    function saveFormData() 
    {
        const formElements = document.querySelectorAll("#twamp-form input");
        formElements.forEach(element => {
            localStorage.setItem(element.id, element.value);
        });
    }

    // Function to load form data from localStorage
    function loadFormData() 
    {
        const formElements = document.querySelectorAll("#twamp-form input");
        formElements.forEach(element => {
            if (localStorage.getItem(element.id)) 
            {
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

    // Clear localStorage when the window or tab is closed
    window.addEventListener("pagehide", function (event) 
    {
        if (!event.persisted) 
        {  
            localStorage.clear();
        }
    });
});

//Field range check
document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('twamp-form');
    const startTestButton = document.getElementById('start-test-button');

    startTestButton.addEventListener('click', () => {
        const packetNumber = document.getElementById('packet-number').value;
        const interval = document.getElementById('interval').value;
        const packetSize = document.getElementById('packet-size').value;
        
        let valid = true;

        if (packetNumber <= 0) 
        {
            alert('Packet number must be greater than 0.');
            valid = false;
        } 
        else if (interval < 10) 
        {
            alert('Interval must be at least 10 ms.');
            valid = false;
        } 
        else if (packetSize < 12 || packetSize > 65507) 
        {
            alert('Packet size must be between 12 and 65507 bytes.');
            valid = false;
        }
        //Prevent from sumbitting if the valid is false
        if (valid) 
        {
            form.submit();
        }
    });
});

