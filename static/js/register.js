document.getElementById('registerForm').addEventListener('submit', (event) => {
    event.preventDefault();
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;

    fetch('http://localhost:8080/user/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email: email, password: password })
    })
    .then(async response => {
        if (!response.ok) {
            const data = await response.json();
            alert(data.error || 'Error al registrarse');
            throw new Error(data.error || 'Error al registrarse');
        }
        else {
            const succesLink = document.getElementById('succesLink');
            succesLink.click();
        }
    })
    .catch(error => {
        console.error('Error:', error);
    });
});