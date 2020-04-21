import { createStyles, Theme, WithStyles } from "@material-ui/core";
import withStyles from "@material-ui/core/styles/withStyles";
import * as React from "react";
import { RouterLink } from "rfluxx-routing";


import { ResourceText } from "../../i18n/Languages";



export const styles = (theme: Theme) => createStyles({
    root: {
    }
});

/**
 * State for { @HomePage }
 */
export interface IHomePageState
{

}

/**
 * Props for { @HomePage }
 */
export interface IHomePageProps
    extends WithStyles<typeof styles>
{
}

/**
 * Component that shows the home page.
 */
export const HomePage = withStyles(styles)(
    class extends React.Component<IHomePageProps, IHomePageState>
    {
        /**
         * Renders the component.
         */
        public render(): React.ReactNode
        {
            const { classes, ...rest } = this.props;

            return <div className={classes.root}>
                
                <ResourceText getText={ x => x.welcome_home } />
                
                <RouterLink path="/log_sources" caption="Log Sources"></RouterLink>
            </div>;
        }
    }
);