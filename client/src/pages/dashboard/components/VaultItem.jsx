import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardAction,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

const VaultItem = ({ vault }) => {
  const navigate = useNavigate();
  const [hovered, setHovered] = useState(false);

  return (
    <Card
      className="w-full border hover:border-primary transition-all relative"
      onMouseEnter={() => setHovered(true)}
      onMouseLeave={() => setHovered(false)}
    >
      <CardHeader>
        <CardTitle className="text-lg">{vault.title}</CardTitle>
        <CardDescription>
          {vault.description ? vault.description : "No description provided."}
        </CardDescription>
        <CardAction>
          <Button
            variant={hovered ? "" : "outline"} // change variant on hover
            className="mt-3 border"
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
