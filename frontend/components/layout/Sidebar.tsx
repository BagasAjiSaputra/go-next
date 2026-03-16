"use client";

import { useState } from "react";
import Link from "next/link";

const navItems = [
    { label: "Dashboard", href: "/dashboard" },
    { label: "Product", href: "/product" },
    { label: "Profile", href: "/profile" },
];

export default function Sidebar() {
    const [collapsed, setCollapsed] = useState(false);

    return (
        <aside
            className={`${
                collapsed ? "w-20" : "w-64"
            } h-screen bg-white border-r border-gray-200 flex flex-col transition-all duration-300 shadow-sm`}
        >
            {/* Header */}
            <div className="flex items-center justify-between px-4 py-5 border-b border-gray-200">
                {!collapsed && <span className="font-bold text-lg text-slate-900">Go Next</span>}
                <button
                    onClick={() => setCollapsed(!collapsed)}
                    className="p-1.5 hover:bg-gray-100 rounded-lg transition-colors"
                    aria-label="Toggle sidebar"
                >
                    {collapsed ? "→" : "←"}
                </button>
            </div>

            {/* Navigation */}
            <nav className="flex-1 px-3 py-4 space-y-1 overflow-y-auto">
                {navItems.map((item) => (
                    <Link
                        key={item.href}
                        href={item.href}
                        className={`flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm text-gray-700 hover:bg-gray-100 transition-colors ${
                            collapsed ? "justify-center" : ""
                        }`}
                        title={collapsed ? item.label : ""}
                    >
                        <span className="w-2 h-2 bg-blue-500 rounded-full flex-shrink-0" />
                        {!collapsed && <span>{item.label}</span>}
                    </Link>
                ))}
            </nav>

            {/* Logout */}
            <div className="px-3 py-4 border-t border-gray-200">
                <form action="Logout">
                    <button
                        type="submit"
                        className={`w-full flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm text-gray-600 hover:bg-red-50 hover:text-red-600 transition-colors ${
                            collapsed ? "justify-center" : ""
                        }`}
                    >
                        <span className="text-lg">↩</span>
                        {!collapsed && <span>Logout</span>}
                    </button>
                </form>
            </div>
        </aside>
    );
}