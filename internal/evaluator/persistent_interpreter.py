import sys
import traceback

global_vars = {}

while True:
    code_lines = []

    while True:
        line = sys.stdin.readline()
        if not line:
            sys.exit(0)

        if line.strip() == '<<<END>>>':
            break

        code_lines.append(line)

    code = ''.join(code_lines)

    try:
        exec_globals = global_vars
        exec_locals = {}
        exec(code, exec_globals, exec_locals)

    except Exception:
        print("<<<EXCEPTION>>>")
        traceback.print_exc(file=sys.stdout)

        print("<<<END_EXCEPTION>>>")
        sys.stdout.flush()

        continue

    print("<<<END_EXECUTION>>>")
    sys.stdout.flush()
