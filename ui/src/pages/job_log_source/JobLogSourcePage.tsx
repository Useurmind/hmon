import { createStyles, Theme, WithStyles } from "@material-ui/core";
import withStyles from "@material-ui/core/styles/withStyles";
import * as React from "react";
import { useCallback } from "react";
import { IUseStoreFromContainerContextProps, useStoreStateFromContainerContext } from "rfluxx-react";

import { JobLogSourceDetails } from "./JobLogSourceDetails";

export const styles = (theme: Theme) => createStyles({
    root: {
    }
});

/**
 * Props for { @JobLogSourcePage }
 */
export interface IJobLogSourcePageProps
    extends WithStyles<typeof styles>
{
}

/**
 * Implementation of JobLogSourcePage.
 */
const JobLogSourcePageImpl: React.FunctionComponent<IJobLogSourcePageProps> = props => {
    const { classes } = props;

    return <div className={classes.root}>
        <JobLogSourceDetails storeRegistrationKey="IJobLogSourceStore"></JobLogSourceDetails>
    </div>;
};

/**
 * Component that shows a job log source.
 */
export const JobLogSourcePage = withStyles(styles)(JobLogSourcePageImpl);
