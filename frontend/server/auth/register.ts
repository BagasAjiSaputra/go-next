'use server'

import { redirect } from "next/navigation"

export async function Register(formData : FormData) {
    
    // const username = formData.get('username')
    const email = formData.get('email')
    const password = formData.get('password')
    
    const res = await fetch('http://localhost:8080/register', {
        method : "POST",
        headers : {"Content-Type" : "application/json"},
        body : JSON.stringify({ email, password})
    })

    const data = await res.json();

    redirect("/login")

}