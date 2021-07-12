package util

import (
	"fmt"
	"testing"
)

func TestRemoveHTML(t *testing.T) {
	txt := `<div id="article-body" class="text-copy bodyCopy auto">
<p>Freeze-dried mouse sperm that spent months at the <a href="https://www.space.com/16748-international-space-station.html">International Space Station</a> (ISS) returned to Earth and successfully fertilized mouse ovary eggs to produce twitchy-nosed "space pups" — for science.&nbsp;</p>
<div id="ad-unit-1" class="ad-unit"></div>
<p>The Japanese researchers behind the new work, which they published today (June 11) in a new paper, wanted to know how space radiation affects fertility in mammals. Radiation can damage the <a href="https://www.livescience.com/37247-dna.html" data-url="https://www.livescience.com/37247-dna.html">DNA</a> within cells, causing mutations (this is why dermatologists recommend using sunscreen). Environments on Earth with heavy radiation exposure can cause defects in the offspring of animals.</p>
<aside class="hawk-placeholder" data-render-type="fte" data-skip="dealsy" data-widget-type="seasonal"></aside>
<p>Space radiation in particular has been a major concern for countries like the U.S. and Japan that have sent many astronauts on lengthy missions into low Earth orbit. Farther space destinations are also on the horizon. NASA and other space agencies are developing systems that could support humans on monthslong journeys to other solar system destinations such as the moon and Mars, and <a href="https://www.space.com/mars-radiation-obstacle-crewed-human-missions.html">radiation is a big concern</a>.&nbsp;</p>
<p><strong>Related: </strong><a href="https://www.space.com/41887-mars-radiation-too-much-for-astronauts.html">Astronauts going to Mars will absorb crazy amounts of radiation</a></p>
<div class="jwplayer__widthsetter">
<div class="jwplayer__wrapper">
<div id="futr_botr_8DCsipf7_bQHItauA_div" class="future__jwplayer">
<div id="botr_8DCsipf7_bQHItauA_div"></div>
</div>
</div>
</div>
<p>This is where the small, squeaky animals enter the story.&nbsp;</p>
<div id="ad-unit-2" class="ad-unit"></div>
<p>Previous studies have been unable to mimic space-radiation conditions on Earth, so this team sent their experiment to space. Researchers freeze-dried mouse sperm samples from 12 mice and sealed them within small lightweight capsules, <a href="https://www.eurekalert.org/jrnls/sciadvances/pages/2021-06/wakayama-06-11-21.php" target="_blank" data-url="https://www.eurekalert.org/jrnls/sciadvances/pages/2021-06/wakayama-06-11-21.php">according to a press release</a> describing the study.&nbsp;</p>
<p>The packets were transported to the ISS and stored for different amounts of time. A portion of the samples returned to Earth after nine months in space, another set returned after two years and nine months, and the final set of mice sperm samples came back after five years and 10 months in space.&nbsp;</p>
<div class="inlinegallery">
<div class="inlinegallery-wrap" style="display:flex; flex-flow:row nowrap;">
<div class="inlinegallery-item" style="flex: 0 0 auto;">
<span class="slidecount">Image 1 of 4</span>
<figure data-bordeaux-image-check>
<div class="image-full-width-wrapper">
<div class="image-widthsetter" style="max-width:2000px;">
<p class="vanilla-image-block" style="padding-top:56.25%;">
<picture>
<source type="image/webp" alt="Mouse sperm were freeze-dried and stored in a glass vial, and stored on the International Space Station for up to five years and 10 months. These samples were later returned to Earth, where researchers checked them for DNA damage." class=" lazy-image-van optional-image" onerror="if(this.src &amp;&amp; this.src.indexOf('missing-image.svg') !== -1){return true;};this.parentNode.replaceChild(window.missingImage(),this)" sizes="(min-width: 1000px) 970px, calc(100vw - 40px)" data-normal="https://vanilla.futurecdn.net/space/media/img/missing-image.svg" data-srcset="https://cdn.mos.cms.futurecdn.net/PEgVeQL8cqMKQBRz5UopAA-320-80.jpg.webp 320w, https://cdn.mos.cms.futurecdn.net/PEgVeQL8cqMKQBRz5UopAA-650-80.jpg.webp 650w, https://cdn.mos.cms.futurecdn.net/PEgVeQL8cqMKQBRz5UopAA-970-80.jpg.webp 970w" data-original-mos="https://cdn.mos.cms.futurecdn.net/PEgVeQL8cqMKQBRz5UopAA.jpg" data-pin-media="https://cdn.mos.cms.futurecdn.net/PEgVeQL8cqMKQBRz5UopAA.jpg">
<source type="image/jpeg" alt="Mouse sperm were freeze-dried and stored in a glass vial, and stored on the International Space Station for up to five years and 10 months. These samples were later returned to Earth, where researchers checked them for DNA damage." class=" lazy-image-van optional-image" onerror="if(this.src &amp;&amp; this.src.indexOf('missing-image.svg') !== -1){return true;};this.parentNode.replaceChild(window.missingImage(),this)" sizes="(min-width: 1000px) 970px, calc(100vw - 40px)" data-normal="https://vanilla.futurecdn.net/space/media/img/missing-image.svg" data-srcset="https://cdn.mos.cms.futurecdn.net/PEgVeQL8cqMKQBRz5UopAA-320-80.jpg 320w, https://cdn.mos.cms.futurecdn.net/PEgVeQL8cqMKQBRz5UopAA-650-80.jpg 650w, https://cdn.mos.cms.futurecdn.net/PEgVeQL8cqMKQBRz5UopAA-970-80.jpg 970w" data-original-mos="https://cdn.mos.cms.futurecdn.net/PEgVeQL8cqMKQBRz5UopAA.jpg" data-pin-media="https://cdn.mos.cms.futurecdn.net/PEgVeQL8cqMKQBRz5UopAA.jpg">
<img src="https://vanilla.futurecdn.net/space/media/img/missing-image.svg" alt="Mouse sperm were freeze-dried and stored in a glass vial, and stored on the International Space Station for up to five years and 10 months. These samples were later returned to Earth, where researchers checked them for DNA damage." class=" lazy-image-van optional-image" onerror="if(this.src &amp;&amp; this.src.indexOf('missing-image.svg') !== -1){return true;};this.parentNode.replaceChild(window.missingImage(),this)" sizes="(min-width: 1000px) 970px, calc(100vw - 40px)" data-normal="https://vanilla.futurecdn.net/space/media/img/missing-image.svg" data-srcset="https://cdn.mos.cms.futurecdn.net/PEgVeQL8cqMKQBRz5UopAA-320-80.jpg 320w, https://cdn.mos.cms.futurecdn.net/PEgVeQL8cqMKQBRz5UopAA-650-80.jpg 650w, https://cdn.mos.cms.futurecdn.net/PEgVeQL8cqMKQBRz5UopAA-970-80.jpg 970w" data-original-mos="https://cdn.mos.cms.futurecdn.net/PEgVeQL8cqMKQBRz5UopAA.jpg" data-pin-media="https://cdn.mos.cms.futurecdn.net/PEgVeQL8cqMKQBRz5UopAA.jpg">
</picture></p>
</div>
</div>
<figcaption itemprop="caption description">
<span class="caption-text">Mouse sperm were freeze-dried and stored in a glass vial, and stored on the International Space Station for up to five years and 10 months. These samples were later returned to Earth, where researchers checked them for DNA damage. </span><span class="credit" itemprop="copyrightHolder">(Image credit: Teruhiko Wakayama/University of Yamanashi)</span>
</figcaption>
</figure>
</div>
<div class="inlinegallery-item" style="flex: 0 0 auto;">
<span class="slidecount">Image 2 of 4</span>
<figure data-bordeaux-image-check>
<div class="image-full-width-wrapper">
<div class="image-widthsetter" style="max-width:1600px;">
<p class="vanilla-image-block" style="padding-top:56.25%;">
<picture>
<source type="image/webp" alt="These mice embryos were fertilized by the freeze-dried mouse sperm that scientists rehydrated with water after they came back from space." class=" lazy-image-van optional-image" onerror="if(this.src &amp;&amp; this.src.indexOf('missing-image.svg') !== -1){return true;};this.parentNode.replaceChild(window.missingImage(),this)" sizes="(min-width: 1000px) 970px, calc(100vw - 40px)" data-normal="https://vanilla.futurecdn.net/space/media/img/missing-image.svg" data-srcset="https://cdn.mos.cms.futurecdn.net/APaQn5EZMne6gkDNSwgnJ7-320-80.jpg.webp 320w, https://cdn.mos.cms.futurecdn.net/APaQn5EZMne6gkDNSwgnJ7-650-80.jpg.webp 650w, https://cdn.mos.cms.futurecdn.net/APaQn5EZMne6gkDNSwgnJ7-970-80.jpg.webp 970w" data-original-mos="https://cdn.mos.cms.futurecdn.net/APaQn5EZMne6gkDNSwgnJ7.jpg" data-pin-media="https://cdn.mos.cms.futurecdn.net/APaQn5EZMne6gkDNSwgnJ7.jpg">
<source type="image/jpeg" alt="These mice embryos were fertilized by the freeze-dried mouse sperm that scientists rehydrated with water after they came back from space." class=" lazy-image-van optional-image" onerror="if(this.src &amp;&amp; this.src.indexOf('missing-image.svg') !== -1){return true;};this.parentNode.replaceChild(window.missingImage(),this)" sizes="(min-width: 1000px) 970px, calc(100vw - 40px)" data-normal="https://vanilla.futurecdn.net/space/media/img/missing-image.svg" data-srcset="https://cdn.mos.cms.futurecdn.net/APaQn5EZMne6gkDNSwgnJ7-320-80.jpg 320w, https://cdn.mos.cms.futurecdn.net/APaQn5EZMne6gkDNSwgnJ7-650-80.jpg 650w, https://cdn.mos.cms.futurecdn.net/APaQn5EZMne6gkDNSwgnJ7-970-80.jpg 970w" data-original-mos="https://cdn.mos.cms.futurecdn.net/APaQn5EZMne6gkDNSwgnJ7.jpg" data-pin-media="https://cdn.mos.cms.futurecdn.net/APaQn5EZMne6gkDNSwgnJ7.jpg">
<img src="https://vanilla.futurecdn.net/space/media/img/missing-image.svg" alt="These mice embryos were fertilized by the freeze-dried mouse sperm that scientists rehydrated with water after they came back from space." class=" lazy-image-van optional-image" onerror="if(this.src &amp;&amp; this.src.indexOf('missing-image.svg') !== -1){return true;};this.parentNode.replaceChild(window.missingImage(),this)" sizes="(min-width: 1000px) 970px, calc(100vw - 40px)" data-normal="https://vanilla.futurecdn.net/space/media/img/missing-image.svg" data-srcset="https://cdn.mos.cms.futurecdn.net/APaQn5EZMne6gkDNSwgnJ7-320-80.jpg 320w, https://cdn.mos.cms.futurecdn.net/APaQn5EZMne6gkDNSwgnJ7-650-80.jpg 650w, https://cdn.mos.cms.futurecdn.net/APaQn5EZMne6gkDNSwgnJ7-970-80.jpg 970w" data-original-mos="https://cdn.mos.cms.futurecdn.net/APaQn5EZMne6gkDNSwgnJ7.jpg" data-pin-media="https://cdn.mos.cms.futurecdn.net/APaQn5EZMne6gkDNSwgnJ7.jpg">
</picture></p>
</div>
</div>
<figcaption itemprop="caption description">
<span class="caption-text">These mice embryos were fertilized by the freeze-dried mouse sperm that scientists rehydrated with water after they came back from space. </span><span class="credit" itemprop="copyrightHolder">(Image credit: Teruhiko Wakayama/University of Yamanashi)</span>
</figcaption>
</figure>
</div>
<div class="inlinegallery-item" style="flex: 0 0 auto;">
<span class="slidecount">Image 3 of 4</span>
<figure data-bordeaux-image-check>
<div class="image-full-width-wrapper">
<div class="image-widthsetter" style="max-width:1600px;">
<p class="vanilla-image-block" style="padding-top:56.25%;">
<picture>
<source type="image/webp" alt="Space sperm were injected into oocytes. This method was called as Intracytoplasmic Sperm Injection (ICSI)." class=" lazy-image-van optional-image" onerror="if(this.src &amp;&amp; this.src.indexOf('missing-image.svg') !== -1){return true;};this.parentNode.replaceChild(window.missingImage(),this)" sizes="(min-width: 1000px) 970px, calc(100vw - 40px)" data-normal="https://vanilla.futurecdn.net/space/media/img/missing-image.svg" data-srcset="https://cdn.mos.cms.futurecdn.net/D3KfxchuQ4L4KyupTzQCj6-320-80.jpg.webp 320w, https://cdn.mos.cms.futurecdn.net/D3KfxchuQ4L4KyupTzQCj6-650-80.jpg.webp 650w, https://cdn.mos.cms.futurecdn.net/D3KfxchuQ4L4KyupTzQCj6-970-80.jpg.webp 970w" data-original-mos="https://cdn.mos.cms.futurecdn.net/D3KfxchuQ4L4KyupTzQCj6.jpg" data-pin-media="https://cdn.mos.cms.futurecdn.net/D3KfxchuQ4L4KyupTzQCj6.jpg">
<source type="image/jpeg" alt="Space sperm were injected into oocytes. This method was called as Intracytoplasmic Sperm Injection (ICSI)." class=" lazy-image-van optional-image" onerror="if(this.src &amp;&amp; this.src.indexOf('missing-image.svg') !== -1){return true;};this.parentNode.replaceChild(window.missingImage(),this)" sizes="(min-width: 1000px) 970px, calc(100vw - 40px)" data-normal="https://vanilla.futurecdn.net/space/media/img/missing-image.svg" data-srcset="https://cdn.mos.cms.futurecdn.net/D3KfxchuQ4L4KyupTzQCj6-320-80.jpg 320w, https://cdn.mos.cms.futurecdn.net/D3KfxchuQ4L4KyupTzQCj6-650-80.jpg 650w, https://cdn.mos.cms.futurecdn.net/D3KfxchuQ4L4KyupTzQCj6-970-80.jpg 970w" data-original-mos="https://cdn.mos.cms.futurecdn.net/D3KfxchuQ4L4KyupTzQCj6.jpg" data-pin-media="https://cdn.mos.cms.futurecdn.net/D3KfxchuQ4L4KyupTzQCj6.jpg">
<img src="https://vanilla.futurecdn.net/space/media/img/missing-image.svg" alt="Space sperm were injected into oocytes. This method was called as Intracytoplasmic Sperm Injection (ICSI)." class=" lazy-image-van optional-image" onerror="if(this.src &amp;&amp; this.src.indexOf('missing-image.svg') !== -1){return true;};this.parentNode.replaceChild(window.missingImage(),this)" sizes="(min-width: 1000px) 970px, calc(100vw - 40px)" data-normal="https://vanilla.futurecdn.net/space/media/img/missing-image.svg" data-srcset="https://cdn.mos.cms.futurecdn.net/D3KfxchuQ4L4KyupTzQCj6-320-80.jpg 320w, https://cdn.mos.cms.futurecdn.net/D3KfxchuQ4L4KyupTzQCj6-650-80.jpg 650w, https://cdn.mos.cms.futurecdn.net/D3KfxchuQ4L4KyupTzQCj6-970-80.jpg 970w" data-original-mos="https://cdn.mos.cms.futurecdn.net/D3KfxchuQ4L4KyupTzQCj6.jpg" data-pin-media="https://cdn.mos.cms.futurecdn.net/D3KfxchuQ4L4KyupTzQCj6.jpg">
</picture></p>
</div>
</div>
<figcaption itemprop="caption description">
<span class="caption-text">Space sperm were injected into oocytes. This method was called as Intracytoplasmic Sperm Injection (ICSI). </span><span class="credit" itemprop="copyrightHolder">(Image credit: Teruhiko Wakayama/University of Yamanashi)</span>
</figcaption>
</figure>
</div>
<div class="inlinegallery-item" style="flex: 0 0 auto;">
<span class="slidecount">Image 4 of 4</span>
<figure data-bordeaux-image-check>
<div class="image-full-width-wrapper">
<div class="image-widthsetter" style="max-width:1600px;">
<p class="vanilla-image-block" style="padding-top:56.25%;">
<picture>
<source type="image/webp" alt="Ampoules were returned from ISS to Earth, then, freeze-dried sperm were rehydrated by water. You can use those sperm for in vitro fertilization immediately, no need to wait 3 minutes." class=" lazy-image-van optional-image" onerror="if(this.src &amp;&amp; this.src.indexOf('missing-image.svg') !== -1){return true;};this.parentNode.replaceChild(window.missingImage(),this)" sizes="(min-width: 1000px) 970px, calc(100vw - 40px)" data-normal="https://vanilla.futurecdn.net/space/media/img/missing-image.svg" data-srcset="https://cdn.mos.cms.futurecdn.net/LznYgS3bXpqpixX4ztuTt8-320-80.jpg.webp 320w, https://cdn.mos.cms.futurecdn.net/LznYgS3bXpqpixX4ztuTt8-650-80.jpg.webp 650w, https://cdn.mos.cms.futurecdn.net/LznYgS3bXpqpixX4ztuTt8-970-80.jpg.webp 970w" data-original-mos="https://cdn.mos.cms.futurecdn.net/LznYgS3bXpqpixX4ztuTt8.jpg" data-pin-media="https://cdn.mos.cms.futurecdn.net/LznYgS3bXpqpixX4ztuTt8.jpg">
<source type="image/jpeg" alt="Ampoules were returned from ISS to Earth, then, freeze-dried sperm were rehydrated by water. You can use those sperm for in vitro fertilization immediately, no need to wait 3 minutes." class=" lazy-image-van optional-image" onerror="if(this.src &amp;&amp; this.src.indexOf('missing-image.svg') !== -1){return true;};this.parentNode.replaceChild(window.missingImage(),this)" sizes="(min-width: 1000px) 970px, calc(100vw - 40px)" data-normal="https://vanilla.futurecdn.net/space/media/img/missing-image.svg" data-srcset="https://cdn.mos.cms.futurecdn.net/LznYgS3bXpqpixX4ztuTt8-320-80.jpg 320w, https://cdn.mos.cms.futurecdn.net/LznYgS3bXpqpixX4ztuTt8-650-80.jpg 650w, https://cdn.mos.cms.futurecdn.net/LznYgS3bXpqpixX4ztuTt8-970-80.jpg 970w" data-original-mos="https://cdn.mos.cms.futurecdn.net/LznYgS3bXpqpixX4ztuTt8.jpg" data-pin-media="https://cdn.mos.cms.futurecdn.net/LznYgS3bXpqpixX4ztuTt8.jpg">
<img src="https://vanilla.futurecdn.net/space/media/img/missing-image.svg" alt="Ampoules were returned from ISS to Earth, then, freeze-dried sperm were rehydrated by water. You can use those sperm for in vitro fertilization immediately, no need to wait 3 minutes." class=" lazy-image-van optional-image" onerror="if(this.src &amp;&amp; this.src.indexOf('missing-image.svg') !== -1){return true;};this.parentNode.replaceChild(window.missingImage(),this)" sizes="(min-width: 1000px) 970px, calc(100vw - 40px)" data-normal="https://vanilla.futurecdn.net/space/media/img/missing-image.svg" data-srcset="https://cdn.mos.cms.futurecdn.net/LznYgS3bXpqpixX4ztuTt8-320-80.jpg 320w, https://cdn.mos.cms.futurecdn.net/LznYgS3bXpqpixX4ztuTt8-650-80.jpg 650w, https://cdn.mos.cms.futurecdn.net/LznYgS3bXpqpixX4ztuTt8-970-80.jpg 970w" data-original-mos="https://cdn.mos.cms.futurecdn.net/LznYgS3bXpqpixX4ztuTt8.jpg" data-pin-media="https://cdn.mos.cms.futurecdn.net/LznYgS3bXpqpixX4ztuTt8.jpg">
</picture></p>
</div>
</div>
<figcaption itemprop="caption description">
<span class="caption-text">Ampoules were returned from the ISS to Earth, then, freeze-dried sperm were rehydrated by water. </span><span class="credit" itemprop="copyrightHolder">(Image credit: Teruhiko Wakayama/University of Yamanashi)</span>
</figcaption>
</figure>
</div>
</div>
</div>
<p>Once back on Earth, the team then determined how much radiation the samples absorbed using RNA sequencing. They found that the ISS trip did not result in DNA damage to the sperm nuclei.&nbsp;</p>
<div id="in-article" class="ad-unit"></div>
<p>They chose to rehydrate the sperm with water, then injected them into fresh mouse ovary cells. After transferring them to female mice, the mothers became pregnant and eventually gave birth to baby mice.&nbsp;</p>
<p>The "space pups" were born healthy and with no defects, according to the team.&nbsp;</p>
<p>The paper detailing the research was published today (June 11) in the journal <a href="https://advances.sciencemag.org/content/7/24/eabg5554" target="_blank" data-url="https://advances.sciencemag.org/content/7/24/eabg5554">Science Advances</a>.&nbsp;</p>
<p><em>Follow Doris Elin Urrutia on Twitter @salazar_elin. Follow us</em> <em>on Twitter @Spacedotcom and on Facebook. </em>&nbsp;</p>
</div>`

	txt = RemoveHTML(txt)

	fmt.Printf("%s \n", txt)
}
