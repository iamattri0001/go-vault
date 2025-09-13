import { Button } from "@/components/ui/button";
import {
  Card,
  CardAction,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import React, { useState } from "react";
import ViewPassword from "./ViewPassword";

const PasswordItem = ({ password, vault_id }) => {
  const [hovered, setHovered] = useState(false);
  return (
    <Card
      className="w-full border hover:border-primary transition-all"
      onMouseEnter={() => setHovered(true)}
      onMouseLeave={() => setHovered(false)}
    >
      <CardHeader>
        <CardTitle className={"text-lg"}>{password.title}</CardTitle>
        <CardDescription>
          {password.description
            ? password.description
            : "No description provided."}
        </CardDescription>
        <CardAction>
          <ViewPassword
            passwordObject={password}
            vault_id={vault_id}
            hovered={hovered}
          />
        </CardAction>
      </CardHeader>
    </Card>
  );
};

export default PasswordItem;
