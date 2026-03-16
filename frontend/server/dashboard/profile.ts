'use server'

import { cookies } from "next/headers";
import { redirect } from "next/navigation";

type Profile = {
  email:    string
  id:       string
  username: string
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