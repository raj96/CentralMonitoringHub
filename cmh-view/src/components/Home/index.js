import { Box } from "@mui/material";
import DataCard from "./DataCard";

export default function Home() {
  return (
    <Box
      style={{
        display: "flex",
        flexDirection: "row",
        justifyContent: "flex-start",
        alignItems: "center",
        gap: "2rem",
        flexWrap: "wrap",
      }}
    >
      <DataCard />
    </Box>
  );
}
