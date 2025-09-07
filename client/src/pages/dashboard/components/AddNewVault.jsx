import { CreateVault } from "@/api/handlers";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { CapitalizeFirstLetter } from "@/utils/text";
import { useState } from "react";
import { toast } from "sonner";

export function AddNewVault({ setVaults }) {
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");

  const handleCreateVault = async () => {
    const response = await CreateVault(title, description);
    if (response.success) {
      toast.success("Vault created successfully!");
      setVaults((prev) => [...prev, response.data.vault]);
    } else {
      if (response.error !== "") {
        toast.error(CapitalizeFirstLetter(response.error));
      } else {
        toast.error("Logout failed. Please try again.");
      }
    }
  };

  return (
    <AlertDialog>
      <AlertDialogTrigger asChild>
        <Button variant="">Create</Button>
      </AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>New Vault</AlertDialogTitle>
          <AlertDialogDescription>
            Give your vault a name and (optionally) a short description. Your
            secrets will be stored securely inside this vault.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <div className="space-y-2">
          <Input
            placeholder="Vault name (e.g. Work, Personal, Banking)"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
          />
          <Textarea
            placeholder="Description (optional)"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
          />
        </div>
        <AlertDialogFooter>
          <AlertDialogCancel>Cancel</AlertDialogCancel>
          <AlertDialogAction onClick={handleCreateVault}>
            Create Vault
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
}
