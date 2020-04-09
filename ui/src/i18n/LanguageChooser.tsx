import { createStyles, MenuItem, Select, Theme, WithStyles } from "@material-ui/core";
import withStyles from "@material-ui/core/styles/withStyles";
import * as React from "react";
import { subscribeStoreSelect, IResolveStoreFromContainerProps } from "rfluxx-react";
import { IResourceStore, IResourceStoreState } from "rfluxx-i18n";
import { usePageContext, IPageContextProps } from "rfluxx-routing";

import { Language, ResourceText } from "./Languages";
import { ResourceTexts } from "./Resources.en";

export const styles = (theme: Theme) => createStyles({
    root: {
    }
});

/**
 * State for { @LanguageChooser }
 */
export interface ILanguageChooserState
{

}

/**
 * Props for { @LanguageChooser }
 */
export interface ILanguageChooserProps
    extends WithStyles<typeof styles>, IPageContextProps
{
    /**
     * The languages that are available for choosing.
     */
    languages: Language[];

    /**
     * The currently selected language.
     */
    activeLanguage: Language;

    /**
     * Change the language.
     */
    setLanguage: (e) => void;
}

/**
 * Component that allows to select a language.
 */
export const LanguageChooserComp = withStyles(styles)(
    class extends React.Component<ILanguageChooserProps, ILanguageChooserState>
    {
        /**
         * Renders the component.
         */
        public render(): React.ReactNode
        {
            const { classes, languages, activeLanguage, setLanguage, ...rest } = this.props;

            return <Select value={activeLanguage ? activeLanguage.key : "" }
                            onChange={setLanguage}>
                {languages && languages.map(x => {
                    return <MenuItem value={x.key}>
                        {x.caption}
                    </MenuItem>;
                })}
            </Select>;
        }
    }
);

// this code binds the component to the ResourceStore
// the actual instance of the ResourceStore is not yet given
const LanguageChooserBound = subscribeStoreSelect<IResourceStore<ResourceTexts>, IResourceStoreState<ResourceTexts>>()(
    LanguageChooserComp,
    (storeState, store) => ({
        // bind the stores state to this components props
        activeLanguage: storeState.activeLanguage,
        languages: storeState.availableLanguages,
        setLanguage: e => store.setLanguage.trigger(e.target.value)
    })
);

// not so nice yet :( but the type system currently fails here
export const LanguageChooser = usePageContext(LanguageChooserBound) as React.ComponentType<IResolveStoreFromContainerProps<IResourceStore<ResourceTexts>, IResourceStoreState<ResourceTexts>>>;