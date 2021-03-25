import { Card, Chip, CardContent, Typography } from "@material-ui/core";
import FaceIcon from '@material-ui/icons/Face';
import DoneIcon from '@material-ui/icons/Done';
import DirectionsRunIcon from '@material-ui/icons/DirectionsRun';

export default function ActiveGroup(props) {
    let participants = []
    if (props.participants) {
        for (let partipant of props.participants) {
        let icon = partipant.raceTimeMs > 0 ? <DoneIcon /> : <DirectionsRunIcon />
        let label = partipant.startNumber + ' '

        if (participants.raceTimeMs > 0) {
            label += partipant.raceTimeMs
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
            {participants}
            {emptyState}
        </CardContent>
    </Card>)
}