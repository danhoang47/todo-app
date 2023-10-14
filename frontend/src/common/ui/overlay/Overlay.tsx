
import "./Overlay.styles.scss"

function Overlay({ children }: { children: React.ReactNode }) {
    return (  
        <div className="overlay" >
            {children}
        </div>
    );
}

export default Overlay;