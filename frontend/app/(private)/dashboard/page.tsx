import { GetProfile } from "@/server/dashboard/profile";

export default async function DashboardPage() {

    const profile = await GetProfile();

    return(
       <div className="min-h-screen bg-gray-50 flex items-center justify-center p-4">
            <div className="bg-white rounded-2xl shadow-sm border border-gray-100 w-full max-w-sm p-8">

                {/* Avatar */}
                <div className="flex flex-col items-center mb-6">
                    <div className="w-16 h-16 rounded-full bg-gray-900 flex items-center justify-center text-white text-xl font-semibold mb-3">
                        {profile?.username?.charAt(0).toUpperCase()}
                    </div>
                    <h1 className="text-lg font-semibold text-gray-900">{profile?.username}</h1>
                </div>

                {/* Info */}
                <div className="space-y-3 mb-6">
                    <div className="flex justify-between items-center py-2 border-b border-gray-50">
                        <span className="text-xs text-gray-400 uppercase tracking-wide">Email</span>
                        <span className="text-sm text-gray-700">{profile?.email}</span>
                    </div>
                    <div className="flex justify-between items-center py-2 border-b border-gray-50">
                        <span className="text-xs text-gray-400 uppercase tracking-wide">ID</span>
                        <span className="text-xs text-gray-400 font-mono truncate max-w-[160px]">{profile?.id}</span>
                    </div>
                </div>

                {/* Logout */}
                <form action={"Logout"}>
                    <button
                        type="submit"
                        className="w-full py-2 text-sm text-gray-500 border border-gray-200 rounded-lg hover:bg-gray-50 hover:text-gray-700 transition-colors"
                    >
                        Logout
                    </button>
                </form>

            </div>
        </div>
    )

}