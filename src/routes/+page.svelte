<script>
	import andrii from '$lib/images/andrii.webp';
	import LinkedIn from '../icons/linkedin.svelte';
	import Instagram from '../icons/instagram.svelte';
	import Github from '../icons/github.svelte';
	import Timeline from '$lib/Timeline.svelte';
	import TimelineItem from '$lib/TimelineItem.svelte';
	import MultiLineMarquee from '$lib/MultiLineMarquee.svelte';
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="def4alt's blog" />
</svelte:head>

<div class="video-background">
	<video src="https://static.def4alt.com/mountain_range.webm" loop muted autoplay playsinline />
	<div class="video-overlay" />
</div>
<section class="profile">
	<div class="profile-cover">
		<div class="profile-name">
			<p>Andrii<br />Olkhovych</p>
			<a
				class="transcription-link"
				href="http://ipa-reader.xyz/?text=%C9%91n%CB%88drij%20%C9%94l%CB%88x%C9%94v%C9%AAt%CA%83&voice=Ewa"
				>/andrij olhovich/</a
			>
		</div>
		<img src={andrii} alt="Andrii Olkhovych" />
	</div>
</section>
<section class="content">
	<aside class="personal-info">
		<p>programmer, student</p>
		<ul class="social-links">
			<a href="https://linkedin.com/in/def4alt"><LinkedIn /></a>
			<a href="https://instagram.com/def4alt"><Instagram /></a>
			<a href="https://github.com/def4alt"><Github /></a>
		</ul>
	</aside>
	<Timeline title="My alma mater">
		<TimelineItem
			content="ukrainian physics and mathematics lyceum"
			link="http://upml.knu.ua/about-upml/"
			startYear={2018}
			endYear={2022}
		/>
		<TimelineItem
			content="taras shevchenko national university of kyiv"
			link="https://knu.ua/en"
			startYear={2022}
			endYear={2023}
		/>
		<TimelineItem
			content="studienkolleg münchen"
			link="https://xn--studienkolleg-mnchen-3ec.de/"
			startYear={2023}
			inProgress={true}
		/>
	</Timeline>
	<div class="work-hobbies-section">
		<div>
			<Timeline title="I worked at">
				<TimelineItem
					content="bamberger haus 🍻"
					startYear={2023}
					link="https://www.bambergerhaus.com/"
				/>
				<TimelineItem
					content="steinheil 16 💁"
					startYear={2023}
					endYear={2024}
					link="https://www.steinheil16.de"
				/>
			</Timeline>
		</div>

		<div>
			<Timeline title="My hobbies">
				<TimelineItem content="traveling ✈️️" />
				<TimelineItem content="swimming 🏊" />
				<TimelineItem content="playing guitar 🎸" />
			</Timeline>
		</div>
	</div>

	<section>
		<h2>I speak</h2>
		<div class="languages">
			<MultiLineMarquee
				title="ukrainian"
				content="Живеш на сьомому поверсі, стіни кімнати завісив козаками й діячами ЗУНР, з вікна бачиш московські дахи, безрадісні тополині алеї, Останкінської телевежі не бачиш — її видно з кімнат, розташованих по другий бік коридору, — але близька її присутність відчувається щохвилини; випромінює щось дуже снодійне, віруси млявості та апатії, тому вранці ніяк не можеш прокинутися, переходиш з одного сновидіння в інше, ніби з країни в країну. Спиш самовіддано, найчастіше до одинадцятої, поки узбек за стіною не увімкне на повний регулятор духмяну орієнтальну музику 'адин палка два струна'. Проклинаючи незлостиво нещасну нашу історію, дружбу народів і союзний договір 22-го року, усвідомлюєш: далі спати не можна."
			/>
			<MultiLineMarquee
				title="english"
				content="It was a bright cold day in April, and the clocks were striking thirteen. Winston Smith, his chin nuzzled into his breast in an effort to escape the vile wind, slipped quickly through the glass doors of Victory Mansions, though not quickly enough to prevent a swirl of gritty dust from entering along with him"
			/>
			<MultiLineMarquee
				title="german"
				content="Ihr naht euch wieder, schwankende Gestalten! Die früh sich einst dem trüben Blick gezeigt. Versuch ich wohl euch diesmal fest zu halten? Fühl ich mein Herz noch jenem Wahn geneigt? Ihr drängt euch zu! nun gut, so mögt ihr walten, 5 Wie ihr aus Dunst und Nebel um mich steigt; Mein Busen fühlt sich jugendlich erschüttert Vom Zauberhauch, der euren Zug umwittert."
			/>
			<MultiLineMarquee
				title="🚧 spanish 🚧"
				bordered={true}
				content="Paracéme, Sancho, que no hay refrán que no sea verdadero, porque todos son sentencias sacadas de la misma experiencia, madre de las ciencias todas."
			/>
		</div>
	</section>
</section>

<style>
	.video-background {
		position: absolute;
		top: 0;
		left: 0;
		width: 100vw;
		height: 100vh;
		min-height: 350px;
		z-index: -1;
		overflow: hidden;
		background-color: var(--color-text);
		background-image: url('https://static.def4alt.com/mountain_range.webp');
		background-repeat: no-repeat;
		background-size: auto 100vh;
		background-position: top;

		& video {
			position: absolute;
			z-index: -2;
			width: 100%;
			height: 100%;
			object-fit: cover;
			overflow: clip;
			overflow-clip-margin: content-box;
			display: inline-block;
			vertical-align: baseline;
		}

		& .video-overlay {
			height: 30vh;
			min-height: 170px;
			width: 100vw;
			position: absolute;
			top: 0;
			left: 0;
			background: linear-gradient(
				180deg,
				rgba(0, 0, 0, 0.86) 0%,
				rgba(0, 0, 0, 0.65) 35%,
				rgba(0, 0, 0, 0) 100%
			);

			@media screen and (orientation: portrait) {
				background: linear-gradient(
					180deg,
					rgba(0, 0, 0, 0.86) 0%,
					rgba(0, 0, 0, 0.65) 55%,
					rgba(0, 0, 0, 0) 100%
				);
			}
		}
	}

	.profile {
		padding-top: var(--navbar-height);
		padding-bottom: 20px;
		height: 100vh;
		min-height: 350px;
		width: 100%;
		display: flex;
		justify-content: center;
		justify-items: center;
		align-items: center;

		overflow: hidden;

		& .profile-cover {
			display: flex;
			justify-content: end;
			align-items: end;

			@media screen and (orientation: portrait) {
				flex-direction: column-reverse;
				justify-content: center;
				align-items: center;
			}

			@media screen and (max-height: 450px) {
				margin-top: -50px;
			}

			& .profile-name {
				display: flex;
				flex-direction: column;
				justify-content: center;
				align-items: start;
				margin-right: -15vw;
				z-index: 3;
				font-size: calc(1rem + 5vw);

				@media screen and (orientation: portrait) {
					display: flex;
					flex-direction: column;
					align-items: center;
					margin-right: 0;
					text-align: center;
					margin-top: 20px;
				}

				@media screen and (min-width: 1440px) {
					margin-right: 0;
				}

				& p {
					font-weight: 700;
					margin: 0;
					line-height: 1.1;
					font-variant-alternates: styleset(alternate);
					color: var(--color-background);
				}

				& .transcription-link {
					font-size: 0.5em;
					color: var(--color-secondary-light);
				}
			}

			& img {
				position: relative;
				object-fit: contain;
				width: auto;
				height: auto;
				border-radius: 15px;
				max-height: 60vh;
				min-width: 300px;
				min-height: 250px;

				@media screen and (orientation: portrait) {
					width: 80%;
					max-height: 40vh;
					min-width: 200px;
					height: 500px;
					object-fit: cover;
					border-radius: 35px;
				}
			}
		}
	}

	.content {
		gap: 50px;
	}

	.personal-info {
		display: flex;
		flex-direction: row;
		gap: 20px;
		justify-content: space-between;

		& p {
			font-size: 1.5rem;

			@media screen and (orientation: portrait) {
				font-size: 2rem;
			}
		}

		& .social-links {
			display: flex;
			padding: 0;
			flex-direction: row;
			align-items: center;
			justify-content: space-between;
			gap: 50px;

			@media screen and (max-width: 768px) {
				gap: 25px;
				flex-direction: column;
			}

			& * {
				width: 40px;
				height: 40px;
				opacity: 90%;
				transition: transform 0.1s ease-in-out;

				@media screen and (max-width: 425px) {
					width: 30px;
					height: 30px;
				}
			}

			& > *:hover {
				opacity: 1;
				transform: scale(1.1);
			}
		}
	}

	.work-hobbies-section {
		display: flex;
		flex-direction: row;

		& > * {
			width: 50%;
		}

		@media screen and (max-width: 768px) {
			gap: 50px;
			flex-direction: column;
			& > * {
				width: 100%;
			}
		}
	}

	.languages {
		width: 100%;
		display: flex;
		flex-direction: column;
		gap: 30px;
		& > * {
			border-radius: 15px;
		}
	}
</style>
