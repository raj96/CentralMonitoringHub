import {
  Drawer,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
} from "@mui/material";
import { Link } from "react-router-dom";
import { v4 as uuid } from "uuid";

export default function SideMenu({
  drawerOpen,
  setDrawerOpen,
  menuContent,
  menuSelectCallback,
  setParentContainerMaxWidth,
}) {
  return (
    <Drawer open={drawerOpen} onClose={() => setDrawerOpen(false)}>
      <List>
        {menuContent.map((content) => (
          <Link key={uuid()} to={content.location} style={{textDecoration: "none", color: "black"}}>
            <ListItem
              button
              onClick={() => {
                menuSelectCallback(content.text);
                setParentContainerMaxWidth(content.parentContainerMaxWidth);
                setDrawerOpen(false);
              }}
            >
              <ListItemIcon>{content.icon}</ListItemIcon>
              <ListItemText>{content.text}</ListItemText>
            </ListItem>
          </Link>
        ))}
      </List>
    </Drawer>
  );
}
