import { ILanguage, CreateResourceTextComponent } from "rfluxx-i18n";
import { ResourceTexts, resources as resourcesEn } from "./Resources.en";
import { resources as resourcesDe } from "./Resources.de";

export const ResourceText = CreateResourceTextComponent<ResourceTexts>();

export type Language = ILanguage<ResourceTexts>;

export const languages: Language[] = [
    {
        key: "en",
        caption: "English",
        resources: resourcesEn
    },
    {
        key: "de",
        caption: "Deutsch",
        resources: resourcesDe
    }
];