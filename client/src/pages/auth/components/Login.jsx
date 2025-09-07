import { LoginUser } from "@/api/handlers";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardAction,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useUserContext } from "@/contexts/UserContext";
import { CapitalizeFirstLetter } from "@/utils/text";
import { Loader2Icon } from "lucide-react";
import { useEffect, useState } from "react";
import { toast } from "sonner";

export function Login({ setAuthType }) {
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    setTimeout(() => setLoading(false), 1000);
  }, []);

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const { login } = useUserContext();

  const handleLogin = async () => {
    if (password.length < 8) {
      toast.error("Password must be at least 8 characters long.");
      return;
    }
    if (username.length < 3 || username.length > 15) {
      toast.error("Username must be between 3 and 15 characters long.");
      return;
    }
    const response = await LoginUser(username, password);
    if (response.success) {
      toast.success("Login successful!");
      login(response.data.user, password);
    } else {
      if (response.error !== "") {
        toast.error(CapitalizeFirstLetter(response.error));
      } else {
        toast.error("Login failed. Please try again.");
      }
    }
  };

  return (
    <Card className="w-full max-w-sm">
      <CardHeader className={loading ? "opacity-0" : ""}>
        <CardTitle>Login to your account</CardTitle>
        <CardDescription>
          Enter your username below to login to your account
        </CardDescription>
        <CardAction>
          <Button
            variant="link"
            className="cursor-pointer"
            onClick={() => setAuthType("register")}
          >
            Register
          </Button>
        </CardAction>
      </CardHeader>
      <CardContent className={loading ? "opacity-0" : ""}>
        <form>
          <div className="flex flex-col gap-6">
            <div className="grid gap-2">
              <Label htmlFor="username">Username</Label>
              <Input
                id="username"
                type="text"
                required
                value={username}
                onChange={(e) => setUsername(e.target.value)}
              />
            </div>
            <div className="grid gap-2">
              <Label htmlFor="password">Password</Label>
              <Input
                id="password"
                type="password"
                required
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </div>
          </div>
        </form>
      </CardContent>
      <CardFooter className={loading ? "opacity-0" : "flex-col gap-2"}>
        <Button type="submit" className="w-full" onClick={handleLogin}>
          Login
        </Button>
      </CardFooter>
      {loading && (
        <Loader2Icon className="absolute inset-0 m-auto animate-spin" />
      )}
    </Card>
  );
}
