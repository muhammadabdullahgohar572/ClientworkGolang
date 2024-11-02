document.getElementById("createUserForm").addEventListener("submit", async function(event) {
    event.preventDefault();

    const Name = document.getElementById("Name").value;
    const Email = document.getElementById("Email").value;
    const Password = document.getElementById("Password").value;
    const Gender = document.getElementById("Gender").value;
    const Company = document.getElementById("Company").value;

    // Use the correct variable names here
    const userData = { 
        Name, 
        Email, 
        Password,  // Corrected the capitalization to match the variable
        Gender, 
        Company
    };

    try {
        const response = await fetch("http://localhost:8080/test", {  // Corrected URL
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(userData),
        });

        const result = await response.json();
        if (response.ok) {
            document.getElementById("responseMessage").textContent = "User created successfully!";
            document.getElementById("createUserForm").reset();
        } else {
            document.getElementById("responseMessage").textContent = "Error creating user: " + result.message;
        }
    } catch (error) {
        document.getElementById("responseMessage").textContent = "Network error. Please try again.";
    }
});
