import { useEffect } from "react";
import PasswordItem from "./PasswordItem";

const PasswordsList = ({ passwords, setPasswords }) => {
  useEffect(() => {
    // Scroll to the top of the list when passwords change
    const list = document.getElementById("passwords-list");
    if (list) {
      list.scrollTop = 0;
    }
  }, [passwords]);
  return (
    <div className="flex-grow flex items-center justify-center">
      {passwords.length === 0 ? (
        <p className="text-center text-gray-500 ">
          No passwords available. Try creating a new password.
        </p>
      ) : (
        <div className="flex flex-col gap-4 w-full md:w-1/2 mt-4">
          {passwords.map((password) => (
            <PasswordItem key={password.id} password={password} />
          ))}
        </div>
      )}
    </div>
  );
};

export default PasswordsList;
