'use client'

import {Login} from "@/server/auth/login";

export default function LoginPage() {

    return(
        <form action={Login}>
            <input type="email" name="email" placeholder="Email"/>
            <input type="text" name="password" placeholder="Password"/>
            <button type="submit">Submit</button>
        </form>
    )

}