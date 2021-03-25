import { Card, Chip, CardContent, Typography } from "@material-ui/core";
import { FaceIcon, DoneIcon, DirectionsRunIcon } from '@material-ui/icons/ExpandMore';

const ActiveGroup = (props) => {
    let participants = []
    if (props.participants) {
        for (let partipant of props.participants) {
        const icon = partipant.raceTimeMs ? <DoneIcon /> : <DirectionsRunIcon />
        participants.push(<Chip
            icon={<FaceIcon />}
            label={participants.startNumber + ' ' + participants.RaceTimeMs}
            color="primary"
            deleteIcon={icon}
          />)
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

export default ActiveGroup