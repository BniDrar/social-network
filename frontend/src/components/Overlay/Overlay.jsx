import styles from "./Overlay.module.css"

const Overlay = ({display}) => {
    return (
        <div className={display ? styles.display : styles.hidden}></div>
    );
}

export default Overlay;