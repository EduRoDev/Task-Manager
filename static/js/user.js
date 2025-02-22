document.getElementById('loginForm').addEventListener('submit', (event) => {
    event.preventDefault();
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;

    fetch('http://localhost:8080/user/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email: email, password: password })
    })
    .then(response => {
        if (!response.ok) {
            return response.json().then(data => {
                alert(data.error);
                throw new Error(data.error);
            });
        }
        return response.json();
    })
    .then(data => {
        console.log('Respuesta completa del servidor:', data);
        
        if (!data.user) {
            throw new Error('Datos de usuario no presentes en la respuesta');
        }
    
        console.log('Usuario logueado:', data.user);
        localStorage.setItem('userId', data.user.id);

        const successLink = document.getElementById('successLink');
        successLink.click();
    })
    .catch(error => {
        console.error('Error:', error);
    });
});

