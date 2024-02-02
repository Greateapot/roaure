import logging
import re
import sys

from requests import Session


HOST = "http://192.168.10.1"
AUTH_KEY = "Basic YWRtaW46YWRtaW4="  #  admin:admin (default password, encoded with base64)
REBOOT_MESSAGE = "The Broadband Router is rebooting."

RESET_ROUTER_URL = f"{HOST}/resetrouter.html"
REBOOT_URL = f"{HOST}/rebootinfo.cgi?sessionKey={{session_key}}"

SESSION_KEY_PATTERN = re.compile(r"var\s+sessionKey='(\d+)';")
LOGGING_LEVEL = logging.DEBUG


def prepare_logger() -> None:
    root = logging.getLogger()
    root.setLevel(LOGGING_LEVEL)

    handler = logging.StreamHandler(sys.stdout)
    handler.setLevel(LOGGING_LEVEL)
    formatter = logging.Formatter(
        "%(asctime)s - %(name)s - %(levelname)s - %(message)s"
    )
    handler.setFormatter(formatter)
    root.addHandler(handler)


def get_session_key(session: Session) -> int:
    logging.debug("Getting current session key...")
    response = session.get(RESET_ROUTER_URL)
    matches = SESSION_KEY_PATTERN.findall(response.text)
    if len(matches) == 0:
        logging.fatal("Failed to get session key.")
    session_key = int(matches[0])
    logging.debug(f"Current session key is {session_key}")
    return session_key


def reset_router(session: Session, session_key: int) -> bool:
    logging.debug(f"Requesting router for rebooting with session key {session_key}")
    response = session.get(REBOOT_URL.format(session_key=session_key))
    return REBOOT_MESSAGE in response.text


def main() -> None:
    prepare_logger()
    logging.info("Starting router reboot...")
    with Session() as session:
        # Оказывается, достаточно только поля авторизации, кт еще и одинаковый при базовом логине:пароле
        session.headers = {"Authorization": AUTH_KEY}
        session_key = get_session_key(session)
        result = reset_router(session, session_key)
        if result:
            logging.info("Router rebooted successfully!")
        else:
            logging.error("Router rebooted unsuccessfully. Check debug logs.")


if __name__ == "__main__":
    main()
