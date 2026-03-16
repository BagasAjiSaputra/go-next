'use server'

import { Timestamp } from "next/dist/server/lib/cache-handlers/types";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";

type Profile = {
  email:    string
  password: string
  id:       string
  username: string
  phone: number
  location: string
  updated_at : Timestamp
}

export async function GetProfile() {

    const jwt = await cookies()
    const token = jwt.get('token')?.value

    if (!token) {
        redirect("/login")
    }

    const res = await fetch('http://localhost:8080/profile', {
        method : 'GET',
        headers : {
            'Authorization' : `Bearer ${token}`,
            'Content-Type' : 'application/json',
        },
    })

    if (!res.ok) {
        redirect('/login')
        return {} as Profile
    }

    return res.json() as Promise<Profile>
}

export async function UpdateProfile(formData : FormData) {

    const email = formData.get("email")
    const password = formData.get("password")
    const username = formData.get("username")
    const phone = Number(formData.get("phone"))
    const location = formData.get("location")

    const cookieStore = await cookies()
    const token = cookieStore.get("token")?.value
    console.log("TOKEN:", token)

    const res = await fetch('http://localhost:8080/users', {
        method : "PUT",
        headers : {
            "Content-Type" : "application/json",
            "Authorization" : `Bearer ${token}`
        },
        body : JSON.stringify({email, password, username, phone, location})
    })

    if (!res.ok) {
    const err = await res.text()
    console.log("API ERROR:", err)
    throw new Error("Update failed")
}

    // const data = await res.json();
    // return data

    redirect('/dashboard')
}