import { Plus, Search } from "lucide-react";

const HogeHogeLayout = () => {
  return (
    <div className="bg-slate-600 min-h-screen">
      <header className="bg-slate-500 p-4">
        <h1 className="text-white text-2xl font-bold">HogeHoge</h1>
      </header>
      <main className="p-4">
        <div className="flex justify-between items-center mb-6">
          <div className="bg-slate-200 rounded-full flex items-center p-2 flex-grow mr-4">
            <Search className="text-gray-500 mr-2" />
            <input
              type="text"
              placeholder="Hinted search text"
              className="bg-transparent outline-none flex-grow"
            />
          </div>
          <button className="bg-pink-200 text-pink-700 px-4 py-2 rounded-full flex items-center">
            <Plus className="mr-1" size={18} />
            Upload
          </button>
        </div>
        <div className="grid grid-cols-3 gap-4">
          {[...Array(6)].map((_, index) => (
            <div
              key={index}
              className={`bg-white rounded-lg aspect-video ${
                index === 4 ? "ring-2 ring-blue-500" : ""
              }`}
            ></div>
          ))}
        </div>
      </main>
    </div>
  );
};

export default HogeHogeLayout;
