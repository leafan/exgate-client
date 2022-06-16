# websocket
import rel
import _thread
import websocket
import json
import time
import threading
import sys

url = "ws://127.0.0.1:8000/ws_proxy_kline?symbol=btcusdt&exchange=fbinance&interval=5m&isProxy=true"
# establish a websocket connection


def on_message(ws, message):
    print(message)


def on_error(ws, error):
    print(error)


def on_close(ws, close_status_code, close_msg):
    print("### closed ###")


def on_open(ws):
    print("Opened connection")


if __name__ == "__main__":
    websocket.enableTrace(True)
    ws = websocket.WebSocketApp(
        "ws://127.0.0.1:8000/ws_proxy_kline?symbol=btcusdt&exchange=fbinance&interval=5m&isProxy=true",
        on_open=on_open,
        on_message=on_message,
        on_error=on_error,
        on_close=on_close)

    ws.run_forever(dispatcher=rel)  # Set dispatcher to automatic reconnection
    rel.signal(2, rel.abort)  # Keyboard Interrupt
    rel.dispatch()
