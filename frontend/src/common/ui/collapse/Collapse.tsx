import { useState, createContext, useCallback, useContext } from 'react'

export type CollapseProps = {
    visible: boolean,
    direction?: 'vertical' | 'horizontal',
    children: React.ReactNode,
    className?: string
}

export type CollapseContextType = {
    onTriggerButtonClick: () => void
}

const CollapseContext = createContext<CollapseContextType | undefined>(undefined) 

function CollapseTrigger({ renderTrigger }: { renderTrigger: (onClick: () => void) => React.ReactNode }) {
    const { onTriggerButtonClick } = useContext(CollapseContext)!

    return (
        <>
            {renderTrigger(onTriggerButtonClick)}
        </>
    )
}

function Collapse({ visible, direction = 'vertical', children, className = '' }: CollapseProps) {
    const [isVisible, setVisible] = useState(visible)
    const expandClassName = direction === 'vertical' ? 'h-0' : 'w-0' 

    const onTriggerButtonClick = useCallback(() => {
        setVisible(prev => !prev)
    }, [])

    return (  
        <CollapseContext.Provider value={{ onTriggerButtonClick }}>
            <div className={`${!isVisible ? expandClassName : ''} transition-all ${className}`}>
                {children}
            </div>
        </CollapseContext.Provider>
    );
}

Collapse.Trigger = CollapseTrigger

export default Collapse;