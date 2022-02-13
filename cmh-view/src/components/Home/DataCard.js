import { ArrowForwardIos } from "@mui/icons-material";
import {
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  CardHeader,
  Divider,
  Grid,
  Typography,
} from "@mui/material";

export default function DataCard({machineId, lastUpdated, numberOfFields}) {
    return <Card style={{ minWidth: "20rem" }}>
    <CardHeader
      title={machineId}
      subheader={`Last updated ${lastUpdated} seconds ago`}
    ></CardHeader>
    <Divider />
    <CardContent>
      <Grid
        container
        direction="column"
        justifyContent="center"
        alignItems="center"
        
      >
        <Grid item>
          <Typography variant="h2">{numberOfFields}</Typography>
        </Grid>
        <Grid item>
          <Typography variant="h6">data fields</Typography>
        </Grid>
      </Grid>
    </CardContent>
    <CardActions>
      <Grid container justifyContent="flex-end">
        <Grid item>
          <Button
            disableElevation
            disableRipple
            endIcon={<ArrowForwardIos />}
          >
            Expand
          </Button>
        </Grid>
      </Grid>
    </CardActions>
  </Card>;
}