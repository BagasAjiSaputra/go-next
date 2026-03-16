// components/Sidebar.tsx
import Link from "next/link";

const navItems = [
    { label: "Dashboard", href: "/dashboard" },
    { label: "Product",   href: "/product" },
    // { label: "Settings",  href: "/dashboard/settings" },
];

export default function Sidebar() {
    return (
        <aside className="fixed top-0 left-0 w-56 h-screen bg-white border-r border-gray-100 flex flex-col z-10">

            {/* Logo */}
            <div className="px-6 py-5 border-b border-gray-100">
                <span className="text-sm font-semibold text-gray-900 tracking-tight">Go Next</span>
            </div>

            {/* Nav */}
            <nav className="flex-1 px-3 py-4 space-y-0.5 overflow-y-auto">
                {navItems.map((item) => (
                    <Link
                        key={item.href}
                        href={item.href}
                        className="flex items-center gap-3 px-3 py-2 rounded-lg text-sm text-gray-500 hover:text-gray-900 hover:bg-gray-50 transition-colors"
                    >
                        {item.label}
                    </Link>
                ))}
            </nav>

            {/* Logout */}
            <div className="px-3 py-4 border-t border-gray-100">
                <form action={"Logout"}>
                    <button
                        type="submit"
                        className="w-full flex items-center gap-3 px-3 py-2 rounded-lg text-sm text-gray-400 hover:text-red-500 hover:bg-red-50 transition-colors text-left"
                    >
                        Logout
                    </button>
                </form>
            </div>

        </aside>
    );
}