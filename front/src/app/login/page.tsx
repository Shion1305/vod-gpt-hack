import { User } from "lucide-react";
import { useState } from "react";

const LoginForm = () => {
  const [stayLoggedIn, setStayLoggedIn] = useState(false);

  return (
    <div className="flex items-center justify-center h-screen bg-slate-600">
      <div className="bg-slate-700 p-8 rounded-lg shadow-lg w-80">
        <div className="flex justify-center mb-6">
          <div className="bg-white rounded-full p-3">
            <User size={48} className="text-slate-700" />
          </div>
        </div>
        <input
          type="text"
          placeholder="USERNAME"
          className="w-full p-2 mb-4 rounded"
        />
        <input
          type="password"
          placeholder="PASSWORD"
          className="w-full p-2 mb-4 rounded"
        />
        <div className="flex items-center justify-between mb-4">
          <label className="flex items-center text-white">
            <input
              type="checkbox"
              checked={stayLoggedIn}
              onChange={() => setStayLoggedIn(!stayLoggedIn)}
              className="mr-2"
            />
            STAY LOGGED IN
          </label>
          <button className="bg-blue-500 text-white px-4 py-2 rounded">
            LOGIN
          </button>
        </div>
      </div>
    </div>
  );
};

export default LoginForm;
