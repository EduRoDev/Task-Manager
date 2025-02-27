
document.getElementById('forgotPasswordForm').addEventListener('submit', (event) => {
    event.preventDefault();
    const email = document.getElementById('email').value;
    fetch('http://localhost:8080/user/ForgotPass', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email: email })
    })
    .then(async response => {
        if (!response.ok) {
            const data = await response.json();
            alert(data.error);
            throw new Error(data.error);
        }
        return response.json();
    })
    .then(data => {
        console.log('Respuesta completa del servidor:', data);
        const successLink = document.getElementById('successLink');
        successLink.click();
    })
    .catch(error => {
        console.error('Error:', error);
    });
});
