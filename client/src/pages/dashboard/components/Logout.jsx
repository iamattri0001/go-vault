import { LogoutUser } from "@/api/handlers";
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
import { useUserContext } from "@/contexts/UserContext";
import { CapitalizeFirstLetter } from "@/utils/text";
import { MdOutlinePowerSettingsNew } from "react-icons/md";
import { toast } from "sonner";

const Logout = () => {
  const { logout } = useUserContext();
  const handleLogout = async () => {
    const response = await LogoutUser();
    if (response.success) {
      toast.success("Logout successful!");
      logout();
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
        <Button variant="outline">
          <MdOutlinePowerSettingsNew />
        </Button>
      </AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Log out of your account?</AlertDialogTitle>
          <AlertDialogDescription>
            You’ll be signed out on this device. You can log back in anytime.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel>Cancel</AlertDialogCancel>
          <AlertDialogAction onClick={handleLogout}>Log out</AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
};

export default Logout;
