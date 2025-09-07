import { Button } from "@/components/ui/button";
import {
  Card,
  CardAction,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import React from "react";
import { useNavigate } from "react-router-dom";

const VaultItem = ({ vault }) => {
  const navigate = useNavigate();
  return (
    <Card className="w-full">
      <CardHeader>
        <CardTitle className={"text-lg"}>{vault.title}</CardTitle>
        <CardDescription>
          {vault.description ? vault.description : "No description provided."}
        </CardDescription>
        <CardAction>
          <Button
            variant="outline"
            className={"mt-3"}
            onClick={() => navigate(`/vault/${vault.id}`)}
          >
            Open
          </Button>
        </CardAction>
      </CardHeader>
    </Card>
  );
};

export default VaultItem;
