import { Button } from "@/components/ui/button";
import {
  Card,
  CardAction,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import React from "react";
import ViewPassword from "./ViewPassword";

const PasswordItem = ({ password }) => {
  return (
    <Card className="w-full">
      <CardHeader>
        <CardTitle className={"text-lg"}>{password.title}</CardTitle>
        <CardDescription>
          {password.description
            ? password.description
            : "No description provided."}
        </CardDescription>
        <CardAction>
          <ViewPassword password={password} />
        </CardAction>
      </CardHeader>
    </Card>
  );
};

export default PasswordItem;
