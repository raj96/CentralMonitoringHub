import "./Navbar.css";
import { AppBar, IconButton, Toolbar, Typography } from "@mui/material";

import { Menu } from "@mui/icons-material";

export default function Navbar({ menuClicked }) {
  return (
    <AppBar className="appbar" position="static">
      <Toolbar>
        <IconButton
          size="large"
          edge="start"
          color="inherit"
          onClick={menuClicked}
        >
          <Menu />
        </IconButton>
        <Typography variant="h5">Central Monitoring Hub</Typography>
      </Toolbar>
    </AppBar>
  );
}
