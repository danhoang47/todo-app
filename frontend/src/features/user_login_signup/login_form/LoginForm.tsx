import { Form } from "@/common/ui";

function LoginForm() {
    return (  
        <Form
            initialValues={{ username: '', password: '' }}
            onSubmit={(values) => {
                console.log(values)
            }}
        >
            {() => (
                <form>

                </form>
            )}
        </Form>
    );
}

export default LoginForm;