import { createStyles, Theme, WithStyles, Typography } from "@material-ui/core";
import withStyles from "@material-ui/core/styles/withStyles";
import * as React from "react";
import { useCallback } from "react";
import { IUseStoreFromContainerContextProps, useStoreStateFromContainerContext } from "rfluxx-react";

import { IJobLogSourceStore, IJobLogSourceStoreState } from './JobLogSourceStore';

export const styles = (theme: Theme) => createStyles({
    root: {
    }
});

/**
 * Props for { @JobLogSourceDetails }
 */
export interface IJobLogSourceDetailsProps
    extends WithStyles<typeof styles>, IUseStoreFromContainerContextProps
{
}

/**
 * Implementation of JobLogSourceDetails.
 */
const JobLogSourceDetailsImpl: React.FunctionComponent<IJobLogSourceDetailsProps> = props => {
    const { classes } = props;

    const [ storeState, store ] = useStoreStateFromContainerContext<IJobLogSourceStore, IJobLogSourceStoreState>(props);
    // const  = useCallback(() => store..trigger(1), [ store ]);

    if (!storeState)
    {
        return null;
    }

    return <div className={classes.root}>
        <Typography>{storeState.jobLogSourceId}</Typography>
    </div>;
};

/**
 * Component that shows all properties of a job log source.
 */
export const JobLogSourceDetails = withStyles(styles)(JobLogSourceDetailsImpl);
