import random
from config import TYPE_NORMALIZATION, TYPE_MAX_RADIUS, DEFAULT_RADIUS


def normalize_type(google_type: str) -> str:
    return TYPE_NORMALIZATION.get(google_type, google_type)


def get_max_radius(node_type: str) -> float:
    return TYPE_MAX_RADIUS.get(node_type, DEFAULT_RADIUS)


def generate_random_color() -> str:
    return "#{:06x}".format(random.randint(0, 0xFFFFFF))