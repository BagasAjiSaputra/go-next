'use client'
import { Register } from "@/server/auth/register";
import Link from "next/link";


export default function RegisterPage() {
  return (
    <main className="min-h-screen bg-white p-4 flex items-center justify-center">
      <section className="flex w-full max-w-5xl overflow-hidden rounded-3xl border border-slate-200 bg-white shadow-xl md:h-[680px]">
        <div className="hidden md:flex w-1/2 items-center justify-center bg-gradient-to-br from-indigo-500 via-violet-500 to-fuchsia-500 p-6 text-white">
        </div>
        <div className="w-full md:w-1/2 p-8 md:p-12">
          <div className="flex items-center gap-2 text-sm font-semibold text-indigo-600 mb-6">
            {/* <span className="inline-block h-2 w-2 rounded-full bg-indigo-600" /> */}
            <span>Xportdity</span>
          </div>
          <h1 className="text-4xl font-extrabold text-slate-900 md:text-5xl leading-tight">
            Register
          </h1>

          <form action={Register} className="mt-8 space-y-4">
            {/* <div>
              <label className="mb-1 block text-xs font-semibold uppercase tracking-[0.16em] text-slate-500">Username</label>
              <input
                type="text"
                name="username"
                required
                placeholder="Your Name"
                className="w-full rounded-md border border-slate-300 px-3 py-2.5 text-sm outline-none transition focus:border-indigo-500 focus:ring-2 focus:ring-indigo-100"
              />
            </div> */}
            <div>
              <label className="mb-1 block text-xs font-semibold uppercase tracking-[0.16em] text-slate-500">Email</label>
              <input
                type="email"
                name="email"
                required
                placeholder="your@gmail.com"
                className="w-full rounded-md border border-slate-300 px-3 py-2.5 text-sm outline-none transition focus:border-indigo-500 focus:ring-2 focus:ring-indigo-100"
              />
            </div>
            <div>
              <label className="mb-1 block text-xs font-semibold uppercase tracking-[0.16em] text-slate-500">Password</label>
              <input
                type="password"
                name="password"
                required
                placeholder="********"
                className="w-full rounded-md border border-slate-300 px-3 py-2.5 text-sm outline-none transition focus:border-indigo-500 focus:ring-2 focus:ring-indigo-100"
              />
            </div>
            <button
              type="submit"
              className="w-full rounded-md bg-indigo-600 px-4 py-2.5 text-sm font-semibold text-white transition hover:bg-indigo-700"
            >
              Register
            </button>
          </form>

          <p className="mt-4 text-center text-xs text-slate-500 md:text-left">
            Already Have Account? <Link href="/login" className="font-semibold text-indigo-600">Login</Link>
          </p>
        </div>
      </section>
    </main>
  )
}
