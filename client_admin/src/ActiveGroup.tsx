import { Card, CardContent, Typography, Box, IconButton } from "@material-ui/core";
import PlayCircleFilledIcon from '@material-ui/icons/PlayCircleFilled';
import { ParticipantDto } from './model';
import Participants from './Participants'
import { startActiveGroup, useAsyncError } from './service';

interface ActiveGroupProps {
    onRefresh?: Function;
    participants?: ParticipantDto[] | undefined
}

export default function ActiveGroup(props: ActiveGroupProps) {
    const throwError = useAsyncError();

    return (<Card elevation = {3}>
        <CardContent>
            <Box display="flex" alignItems="center">
                <Typography variant="h5" component="h2">Active group</Typography>
                <IconButton size="medium" color="secondary" onClick={() => {
                    startActiveGroup()
                        .then(() => { if (props.onRefresh) {props.onRefresh()}})
                        .catch((err) => throwError(err));
                    }}>
                    <PlayCircleFilledIcon />
                </IconButton>
            </Box>
            <Participants participants={props.participants}></Participants>            
        </CardContent>
    </Card>)
}