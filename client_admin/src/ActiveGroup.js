import { Card, Chip, CardContent, Typography, Button, Divider } from "@material-ui/core";
import FaceIcon from '@material-ui/icons/Face';
import DoneIcon from '@material-ui/icons/Done';
import DirectionsRunIcon from '@material-ui/icons/DirectionsRun';

function startActiveGroup() {
    fetch("http://localhost:8010/groups/active", {
        method: 'POST',
        body: JSON.stringify({dummy: 'dummy'})
    });
}

export default function ActiveGroup(props) {
    let participants = []
    if (props.participants) {
        for (let partipant of props.participants) {
        let icon = partipant.raceTimeMs > 0 ? <DoneIcon /> : <DirectionsRunIcon />
        let label = `#${partipant.startNumber}`

        if (partipant.raceTimeMs > 0) {
            label += ` (${partipant.raceTimeMs / 1000} s)`
        }
        participants.push(<Chip
            icon={icon}
            label={label}
            color="primary"
            deleteIcon={icon}></Chip>
          )
        }
    }
    
    const emptyState = participants.length > 0 ? '' : 'No active group.'


    return (<Card>
        <CardContent>
            <Typography variant="h5" component="h2">Active group</Typography>
            <Button variant="contained"  onClick={() => startActiveGroup()} color="primary">Start</Button>
            <Divider />
            {participants}
            {emptyState}
        </CardContent>
    </Card>)
}