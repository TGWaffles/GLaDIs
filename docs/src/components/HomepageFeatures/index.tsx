import type { ReactNode } from "react";
import clsx from "clsx";
import Heading from "@theme/Heading";
import styles from "./styles.module.css";

type FeatureItem = {
    title: string;
    Svg: React.ComponentType<React.ComponentProps<"svg">>;
    description: ReactNode;
};

const FeatureList: FeatureItem[] = [
    {
        title: "Easy to Use",
        Svg: require("@site/static/img/undraw_programmer_raqr.svg").default,
        description: (
            <>
                Dapper Go was designed to be lightweight and easy to use. By
                design, it allows you to build Discord applications with ease.
            </>
        ),
    },
    {
        title: "(Soon to be) Production Ready",
        Svg: require("@site/static/img/undraw_server_9eix.svg").default,
        description: (
            <>
                Used in many production applications and based of years of
                experience, Dapper Go is a reliable and stable choice for your
                next project.
            </>
        ),
    },
    {
        title: "Get Started Quickly",
        Svg: require("@site/static/img/undraw_chat-bot_44el.svg").default,
        description: (
            <>
                Just focus on your application and let Dapper Go handle the
                rest.
            </>
        ),
    },
];

function Feature({ title, Svg, description }: FeatureItem) {
    return (
        <div className={clsx("col col--4")}>
            <div className="text--center">
                <Svg className={styles.featureSvg} role="img" />
            </div>
            <div className="text--center padding-horiz--md">
                <Heading as="h3">{title}</Heading>
                <p>{description}</p>
            </div>
        </div>
    );
}

export default function HomepageFeatures(): ReactNode {
    return (
        <section className={styles.features}>
            <div className="container">
                <div className="row">
                    {FeatureList.map((props, idx) => (
                        <Feature key={idx} {...props} />
                    ))}
                </div>
            </div>
        </section>
    );
}
