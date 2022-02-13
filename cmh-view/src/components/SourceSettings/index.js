import { Paper, Tab, Tabs } from "@mui/material";
import { useState } from "react";
import SourcePage from "./SourcePage";
import SourceTypePage from "./SourceTypePage";

export default function SourceSettings() {
  const TabValue = {
    SOURCE: 0,
    SOURCETYPE: 1,
  };
  const TabPage = (tabValue) => {
    switch (tabValue) {
      case TabValue.SOURCE:
        return <SourcePage />;
      case TabValue.SOURCETYPE:
        return <SourceTypePage />;
      default:
        return undefined;
    }
  };

  const [selectedTab, setSelectedTab] = useState(TabValue.SOURCE);

  return (
    <>
      <Paper>
        <Tabs value={selectedTab}>
          <Tab
            label="Source List"
            value={TabValue.SOURCE}
            onClick={() => setSelectedTab(TabValue.SOURCE)}
          />
          <Tab
            label="Sourcetype List"
            value={TabValue.SOURCETYPE}
            onClick={() => setSelectedTab(TabValue.SOURCETYPE)}
          />
        </Tabs>
        {TabPage(selectedTab)}
      </Paper>
    </>
  );
}
