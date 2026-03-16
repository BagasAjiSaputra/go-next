// 'use client'

import { UpdateProfile, GetProfile } from "@/server/dashboard/profile";
import { Button } from "@/components/ui/button";

export default async function ProfilePage() {

    const profile = await GetProfile();

    return (
        <>
            <form action={UpdateProfile} className="flex flex-col">
                <h1 className="mb-4 font-semibold">Update Profile</h1>
                <input className="mb-2 w-1/3 bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow" 
                placeholder="Email" 
                name="email"
                type="text"
                defaultValue={profile.email}
                />
                <input className="w-1/3 bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow" 
                placeholder="Password" 
                name="password"
                type="password"
                defaultValue={profile?.password}
                />
                <input className="mb-2 w-1/3 bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow" 
                placeholder="Username" 
                name="username"
                type="text"
                defaultValue={profile.username}
                />
                <input className="mb-2 w-1/3 bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow" 
                placeholder="Phone" 
                name="phone"
                type="number"
                defaultValue={profile.phone}
                />
                <input className="mb-2 w-1/3 bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow" 
                placeholder="Location" 
                name="location"
                type="text"
                defaultValue={profile.location}
                />
                <p className="text-sm">Updated at : {profile.updated_at}</p>
                {/* <button type="submit" className=" w-1/5 text-body bg-neutral-secondary-medium box-border border border-default-medium hover:bg-neutral-tertiary-medium hover:text-heading focus:ring-4 focus:ring-neutral-tertiary shadow-xs font-medium leading-5 rounded-base text-sm px-4 py-2.5 focus:outline-none">Update</button> */}
                <Button type="submit">Update</Button>
            </form>
        </>
    )

}