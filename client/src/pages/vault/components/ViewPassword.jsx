import { useState, useEffect } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { BiShow } from "react-icons/bi";
import { FaEdit } from "react-icons/fa";
import { toast } from "sonner";
import { Toggle } from "@/components/ui/toggle";

import {
  AlertDialog,
  AlertDialogTrigger,
  AlertDialogContent,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogCancel,
  AlertDialogAction,
} from "@/components/ui/alert-dialog";

import { useUserContext } from "@/contexts/UserContext";
import { DecryptString } from "@/utils/encryption";
import { CreatePassword, UpdatePassword } from "@/api/handlers";

const ViewPassword = ({ passwordObject, vault_id, setPasswords, hovered }) => {
  const { encKey } = useUserContext();

  const [editable, setEditable] = useState(false);
  const [isPasswordVisible, setIsPasswordVisible] = useState(false);

  // read-only mode
  const [decryptedPassword, setDecryptedPassword] = useState("");

  // editable mode
  const [title, setTitle] = useState(passwordObject?.title || "");
  const [description, setDescription] = useState(
    passwordObject?.description || ""
  );
  const [username, setUsername] = useState(passwordObject?.username || "");
  const [password, setPassword] = useState(""); // editable password state
  const [website, setWebsite] = useState(passwordObject?.website || "");

  useEffect(() => {
    if (editable) {
      // decrypt password only for editable mode
      (async () => {
        if (!encKey) return;
        try {
          const decrypted = await DecryptString(
            passwordObject.encrypted_password,
            encKey
          );
          setPassword(decrypted); // fill editable field
        } catch {
          toast.error("Failed to decrypt password for editing.");
        }
      })();
    } else {
      setPassword(""); // reset editable password when leaving edit mode
      setIsPasswordVisible(false);
      setDecryptedPassword(""); // reset read-only decrypted password
    }
  }, [editable, encKey, passwordObject]);

  const handleDecryptReadOnly = async () => {
    if (!encKey) {
      toast.error("Encryption key missing. Please unlock your vault.");
      return;
    }
    try {
      const decrypted = await DecryptString(
        passwordObject.encrypted_password,
        encKey
      );
      setDecryptedPassword(decrypted);
      setIsPasswordVisible(true);
    } catch {
      toast.error("Failed to decrypt password.");
    }
  };

  const handleUpdate = async () => {
    if (!encKey) {
      toast.error("Encryption key missing. Please unlock your vault.");
      return;
    }
    const response = await UpdatePassword({
      id: passwordObject.id,
      title,
      description,
      username,
      password,
      website,
      vault_id,
      encryptionKey: encKey,
    });
    if (response.success) {
      toast.success("Password updated successfully!");
      setPasswords((prev) =>
        prev.map((p) =>
          p.id === passwordObject.id ? response.data.password : p
        )
      );
      setEditable(false);
    } else {
      toast.error(response.error || "Failed to update password.");
    }
  };

  const toggleMode = () => {
    setEditable((prev) => {
      if (prev) {
        toast.info("Switched to read-only mode. Any unsaved changes are lost.");
      } else {
        toast.info("Switched to edit mode. You can now update the entry.");
      }
      return !prev;
    });
  };

  const displayedPassword = editable
    ? password
    : decryptedPassword || passwordObject.encrypted_password;

  return (
    <AlertDialog>
      <AlertDialogTrigger asChild>
        <Button variant={hovered ? "" : "outline"} className={"mt-3 border"}>
          View
        </Button>
      </AlertDialogTrigger>

      <AlertDialogContent
        className={
          "sm:max-w-md border" + (editable ? " border-destructive" : "")
        }
      >
        <AlertDialogHeader>
          <AlertDialogTitle className="flex justify-between items-center">
            {editable ? "Edit Vault Entry" : "View Vault Entry"}
            <Button variant={editable ? "" : "outline"} onClick={toggleMode}>
              <FaEdit />
            </Button>
          </AlertDialogTitle>
          <AlertDialogDescription>
            {editable
              ? "Update your saved login entry securely. Changes will be encrypted."
              : "You can review your saved login entry. Decrypt the password to view it."}
          </AlertDialogDescription>
        </AlertDialogHeader>

        <div className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-muted-foreground">
              Vault Entry Name
            </label>
            <Input
              value={title}
              onChange={(e) => editable && setTitle(e.target.value)}
              readOnly={!editable}
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-muted-foreground">
              Notes (optional)
            </label>
            <Textarea
              value={description}
              onChange={(e) => editable && setDescription(e.target.value)}
              readOnly={!editable}
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-muted-foreground">
              Username / Email
            </label>
            <Input
              value={username}
              onChange={(e) => editable && setUsername(e.target.value)}
              readOnly={!editable}
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-muted-foreground">
              Password
            </label>
            <div className="flex gap-x-2">
              <Input
                type={isPasswordVisible ? "text" : "password"}
                value={displayedPassword}
                onChange={(e) => editable && setPassword(e.target.value)}
                readOnly={!editable}
                className={
                  !editable && isPasswordVisible
                    ? "border-destructive font-semibold"
                    : ""
                }
              />
              <Toggle
                onClick={() => {
                  if (!editable) {
                    if (!isPasswordVisible) handleDecryptReadOnly();
                    else {
                      setDecryptedPassword("");
                      setIsPasswordVisible(false);
                    }
                  } else {
                    setIsPasswordVisible((prev) => !prev);
                  }
                }}
              >
                <BiShow />
              </Toggle>
            </div>
          </div>

          <div>
            <label className="block text-sm font-medium text-muted-foreground">
              Website / App
            </label>
            <Input
              value={website}
              onChange={(e) => editable && setWebsite(e.target.value)}
              readOnly={!editable}
            />
          </div>
        </div>

        <AlertDialogFooter>
          <AlertDialogCancel onClick={() => setEditable(false)}>
            Close
          </AlertDialogCancel>
          {editable && (
            <AlertDialogAction onClick={handleUpdate}>Save</AlertDialogAction>
          )}
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
};

export default ViewPassword;
