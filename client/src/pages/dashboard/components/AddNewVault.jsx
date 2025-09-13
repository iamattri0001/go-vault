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

  const resetFields = () => {
    setTitle("");
    setDescription("");
  };

  const handleCreateVault = async () => {
    const response = await CreateVault(title, description);
    if (response.success) {
      toast.success("Vault created successfully!");
      setVaults((prev) => [...prev, response.data.vault]);
      resetFields();
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
            onChange={(e) => {
              if (e.target.value.length <= 30) {
                setTitle(e.target.value);
              } else {
                toast.error("Name cannot exceed 30 characters.");
              }
            }}
          />
          <Textarea
            placeholder="Description (optional)"
            value={description}
            onChange={(e) => {
              if (e.target.value.length <= 100) {
                setDescription(e.target.value);
              } else {
                toast.error("Description cannot exceed 100 characters.");
              }
            }}
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
